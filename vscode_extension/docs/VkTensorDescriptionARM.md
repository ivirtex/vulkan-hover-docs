# VkTensorDescriptionARM(3) Manual Page

## Name

VkTensorDescriptionARM - Structure describing a tensor



## [](#_c_specification)C Specification

The `VkTensorDescriptionARM` structure is defined as:

```c++
// Provided by VK_ARM_tensors
typedef struct VkTensorDescriptionARM {
    VkStructureType          sType;
    const void*              pNext;
    VkTensorTilingARM        tiling;
    VkFormat                 format;
    uint32_t                 dimensionCount;
    const int64_t*           pDimensions;
    const int64_t*           pStrides;
    VkTensorUsageFlagsARM    usage;
} VkTensorDescriptionARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `tiling` is a [VkTensorTilingARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorTilingARM.html) value specifying the tiling of the tensor
- `format` is a one component [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) describing the format and type of the data elements that will be contained in the tensor.
- `dimensionCount` is the number of dimensions for the tensor.
- `pDimensions` is a pointer to an array of integers of size `dimensionCount` providing the number of data elements in each dimension.
- `pStrides` is either `NULL` or is an array of size `dimensionCount` providing the strides in bytes for the tensor in each dimension.
- `usage` is a bitmask of [VkTensorUsageFlagBitsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorUsageFlagBitsARM.html) specifying the usage of the tensor.

## [](#_description)Description

When describing a tensor created with `VK_TENSOR_TILING_OPTIMAL_ARM`, `pStrides` must be equal to `NULL`. When describing a tensor created with `VK_TENSOR_TILING_LINEAR_ARM`, `pStrides` is either an array of size `dimensionCount` or `NULL`.

The formats that **must** be supported for `format` are documented in [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-formats-mandatory-features-tensor](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-formats-mandatory-features-tensor).

Each element in the `pStrides` array describes the offset in bytes between increments of the given dimension. For example, `pStrides`\[0] describes the offset between element \[x0,x1,x2,x3] and element \[x0+1,x1,x2,x3]. The `pStrides` array **can** be used to determine whether a tensor is *packed* or not. If `pStrides`\[`dimensionCount`-1] is equal to the size of a tensor element and for each dimension `n` greater than 0 and less than `dimensionCount`, `pStrides`\[n-1] is equal to `pStrides`\[n] * `pDimensions`\[n], then the tensor is a packed tensor. If the [tensorNonPacked](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-tensorNonPacked) feature is not enabled, the tensor **must** be a packed tensor.

When a tensor is created with `VK_TENSOR_TILING_LINEAR_ARM` and `pStrides` equal to `NULL` the tensor strides are calculated by the vulkan implementation such that the resulting tensor is a packed tensor.

Expressed as an addressing formula, the starting byte of an element in a 4-dimensional, for example, linear tensor has address:

```c
// Assume (x0,x1,x2,x3) are in units of elements.

address(x0,x1,x2,x3) = x0*pStrides[0] + x1*pStrides[1] + x2*pStrides[2] + x3*pStrides[3]
```

Valid Usage

- [](#VUID-VkTensorDescriptionARM-dimensionCount-09733)VUID-VkTensorDescriptionARM-dimensionCount-09733  
  `dimensionCount` **must** be less than or equal to [VkPhysicalDeviceTensorPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTensorPropertiesARM.html)::`maxTensorDimensionCount`
- [](#VUID-VkTensorDescriptionARM-pDimensions-09734)VUID-VkTensorDescriptionARM-pDimensions-09734  
  For each i where i ≤ dimensionCount-1, `pDimensions`\[i] **must** be greater than `0`
- [](#VUID-VkTensorDescriptionARM-pDimensions-09883)VUID-VkTensorDescriptionARM-pDimensions-09883  
  For each i where i ≤ dimensionCount-1, `pDimensions`\[i] **must** be less than or equal to [`VkPhysicalDeviceTensorPropertiesARM`::`maxPerDimensionTensorElements`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxPerDimensionTensorElements)
- [](#VUID-VkTensorDescriptionARM-format-09735)VUID-VkTensorDescriptionARM-format-09735  
  `format` **must** not be `VK_FORMAT_UNDEFINED` and **must** be a one-component [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html)
- [](#VUID-VkTensorDescriptionARM-pStrides-09736)VUID-VkTensorDescriptionARM-pStrides-09736  
  `pStrides`\[`dimensionCount`-1] **must** equal the size in bytes of a tensor element
- [](#VUID-VkTensorDescriptionARM-pStrides-09737)VUID-VkTensorDescriptionARM-pStrides-09737  
  For each i, `pStrides`\[i] **must** be a multiple of the element size
- [](#VUID-VkTensorDescriptionARM-pStrides-09738)VUID-VkTensorDescriptionARM-pStrides-09738  
  For each i, `pStrides`\[i] **must** be greater than `0` and less than or equal to [VkPhysicalDeviceTensorPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTensorPropertiesARM.html)::`maxTensorStride`
- [](#VUID-VkTensorDescriptionARM-pStrides-09884)VUID-VkTensorDescriptionARM-pStrides-09884  
  `pStrides`\[0] × `pDimensions`\[0] **must** be less than or equal to [`VkPhysicalDeviceTensorPropertiesARM`::`maxTensorSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxTensorSize)
- [](#VUID-VkTensorDescriptionARM-pStrides-09739)VUID-VkTensorDescriptionARM-pStrides-09739  
  For each i greater than 0, `pStrides`\[i-1] **must** be greater than or equal to `pStrides`\[i] × `pDimensions`\[i] so that no two elements of the tensor reference the same memory address
- [](#VUID-VkTensorDescriptionARM-None-09740)VUID-VkTensorDescriptionARM-None-09740  
  If the [tensorNonPacked](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-tensorNonPacked) feature is not enabled, then the members of [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html) **must** describe a packed tensor
- [](#VUID-VkTensorDescriptionARM-tiling-09741)VUID-VkTensorDescriptionARM-tiling-09741  
  If `tiling` is `VK_TENSOR_TILING_OPTIMAL_ARM` and `usage` is `VK_TENSOR_USAGE_IMAGE_ALIASING_BIT_ARM` then the size of the tensor along its innermost dimension, i.e. `pDimensions`\[`dimensionCount` - 1], **must** be less than or equal to `4`
- [](#VUID-VkTensorDescriptionARM-tiling-09742)VUID-VkTensorDescriptionARM-tiling-09742  
  If `tiling` is `VK_TENSOR_TILING_LINEAR_ARM` then `VK_TENSOR_USAGE_IMAGE_ALIASING_BIT_ARM` **must** not be set in `usage`

Valid Usage (Implicit)

- [](#VUID-VkTensorDescriptionARM-sType-sType)VUID-VkTensorDescriptionARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_TENSOR_DESCRIPTION_ARM`
- [](#VUID-VkTensorDescriptionARM-tiling-parameter)VUID-VkTensorDescriptionARM-tiling-parameter  
  `tiling` **must** be a valid [VkTensorTilingARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorTilingARM.html) value
- [](#VUID-VkTensorDescriptionARM-format-parameter)VUID-VkTensorDescriptionARM-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value
- [](#VUID-VkTensorDescriptionARM-pDimensions-parameter)VUID-VkTensorDescriptionARM-pDimensions-parameter  
  `pDimensions` **must** be a valid pointer to an array of `dimensionCount` `int64_t` values
- [](#VUID-VkTensorDescriptionARM-pStrides-parameter)VUID-VkTensorDescriptionARM-pStrides-parameter  
  If `pStrides` is not `NULL`, `pStrides` **must** be a valid pointer to an array of `dimensionCount` `int64_t` values
- [](#VUID-VkTensorDescriptionARM-usage-parameter)VUID-VkTensorDescriptionARM-usage-parameter  
  `usage` **must** be a valid combination of [VkTensorUsageFlagBitsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorUsageFlagBitsARM.html) values
- [](#VUID-VkTensorDescriptionARM-usage-requiredbitmask)VUID-VkTensorDescriptionARM-usage-requiredbitmask  
  `usage` **must** not be `0`
- [](#VUID-VkTensorDescriptionARM-dimensionCount-arraylength)VUID-VkTensorDescriptionARM-dimensionCount-arraylength  
  `dimensionCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkPhysicalDeviceExternalTensorInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalTensorInfoARM.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateInfoARM.html), [VkTensorTilingARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorTilingARM.html), [VkTensorUsageFlagsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorUsageFlagsARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTensorDescriptionARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0