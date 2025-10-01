# VkCopyTensorInfoARM(3) Manual Page

## Name

VkCopyTensorInfoARM - Structure specifying an tensor copy operation



## [](#_c_specification)C Specification

The `VkCopyTensorInfoARM` structure is defined as:

```c++
// Provided by VK_ARM_tensors
typedef struct VkCopyTensorInfoARM {
    VkStructureType           sType;
    const void*               pNext;
    VkTensorARM               srcTensor;
    VkTensorARM               dstTensor;
    uint32_t                  regionCount;
    const VkTensorCopyARM*    pRegions;
} VkCopyTensorInfoARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is NULL or a pointer to a structure extending this structure.
- `srcTensor` is the source tensor.
- `dstTensor` is the destination tensor.
- `regionCount` is the number of regions to copy.
- `pRegions` is a pointer to an array of [VkTensorCopyARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCopyARM.html) structures specifying the regions to copy.

## [](#_description)Description

Each region in `pRegions` describes a region to be copied from the source tensor to a corresponding region of the destination tensor. `srcTensor` and `dstTensor` **can** be the same tensor or alias the same memory.

The formats of `srcTensor` and `dstTensor` **must** be compatible. Formats are compatible if they share the same class, as shown in the [Compatible Formats](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-compatibility) table.

`vkCmdCopyTensorARM` allows copying between *size-compatible* internal formats.

Valid Usage

- [](#VUID-VkCopyTensorInfoARM-dimensionCount-09684)VUID-VkCopyTensorInfoARM-dimensionCount-09684  
  `srcTensor` and `dstTensor` **must** have been created with equal values for [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html)::`dimensionCount`
- [](#VUID-VkCopyTensorInfoARM-pDimensions-09685)VUID-VkCopyTensorInfoARM-pDimensions-09685  
  For each of the elements of [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html)::`pDimensions`, `srcTensor` and `dstTensor` **must** be the same size
- [](#VUID-VkCopyTensorInfoARM-regionCount-09686)VUID-VkCopyTensorInfoARM-regionCount-09686  
  `regionCount` must be equal to 1
- [](#VUID-VkCopyTensorInfoARM-pRegions-09687)VUID-VkCopyTensorInfoARM-pRegions-09687  
  Each element of `pRegions` **must** be a [VkTensorCopyARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCopyARM.html) structure whose `pSrcOffset` is `NULL` or has all its elements equal to `0`
- [](#VUID-VkCopyTensorInfoARM-pRegions-09688)VUID-VkCopyTensorInfoARM-pRegions-09688  
  Each element of `pRegions` **must** be a [VkTensorCopyARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCopyARM.html) structure whose `pDstOffset` is `NULL` or has all its elements equal to `0`
- [](#VUID-VkCopyTensorInfoARM-pRegions-09689)VUID-VkCopyTensorInfoARM-pRegions-09689  
  Each element of `pRegions` **must** be a [VkTensorCopyARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCopyARM.html) structure whose `pExtent` is `NULL` or equal to the [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html)::`pDimensions` array specified when `srcTensor` and `dstTensor` were created
- [](#VUID-VkCopyTensorInfoARM-pRegions-09954)VUID-VkCopyTensorInfoARM-pRegions-09954  
  Each element of `pRegions` **must** be a [VkTensorCopyARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCopyARM.html) structure whose `dimensionCount`, if it is not equal to 0, is equal to the largest of the [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html)::`dimensionCount` of `srcTensor` or `dstTensor`
- [](#VUID-VkCopyTensorInfoARM-srcTensor-09690)VUID-VkCopyTensorInfoARM-srcTensor-09690  
  The [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-tensor-view-format-features) of `srcTensor` **must** contain `VK_FORMAT_FEATURE_2_TRANSFER_SRC_BIT`
- [](#VUID-VkCopyTensorInfoARM-srcTensor-09691)VUID-VkCopyTensorInfoARM-srcTensor-09691  
  `srcTensor` **must** have been created with `VK_TENSOR_USAGE_TRANSFER_SRC_BIT_ARM` usage flag
- [](#VUID-VkCopyTensorInfoARM-dstTensor-09692)VUID-VkCopyTensorInfoARM-dstTensor-09692  
  The [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-tensor-view-format-features) of `dstTensor` **must** contain `VK_FORMAT_FEATURE_2_TRANSFER_DST_BIT`
- [](#VUID-VkCopyTensorInfoARM-dstTensor-09693)VUID-VkCopyTensorInfoARM-dstTensor-09693  
  `dstTensor` **must** have been created with `VK_TENSOR_USAGE_TRANSFER_DST_BIT_ARM` usage flag
- [](#VUID-VkCopyTensorInfoARM-srcTensor-09694)VUID-VkCopyTensorInfoARM-srcTensor-09694  
  If `srcTensor` is non-sparse then it **must** be bound completely and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object
- [](#VUID-VkCopyTensorInfoARM-dstTensor-09695)VUID-VkCopyTensorInfoARM-dstTensor-09695  
  If `dstTensor` is non-sparse then it **must** be bound completely and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object

Valid Usage (Implicit)

- [](#VUID-VkCopyTensorInfoARM-sType-sType)VUID-VkCopyTensorInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COPY_TENSOR_INFO_ARM`
- [](#VUID-VkCopyTensorInfoARM-pNext-pNext)VUID-VkCopyTensorInfoARM-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkCopyTensorInfoARM-srcTensor-parameter)VUID-VkCopyTensorInfoARM-srcTensor-parameter  
  `srcTensor` **must** be a valid [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html) handle
- [](#VUID-VkCopyTensorInfoARM-dstTensor-parameter)VUID-VkCopyTensorInfoARM-dstTensor-parameter  
  `dstTensor` **must** be a valid [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html) handle
- [](#VUID-VkCopyTensorInfoARM-pRegions-parameter)VUID-VkCopyTensorInfoARM-pRegions-parameter  
  `pRegions` **must** be a valid pointer to an array of `regionCount` valid [VkTensorCopyARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCopyARM.html) structures
- [](#VUID-VkCopyTensorInfoARM-regionCount-arraylength)VUID-VkCopyTensorInfoARM-regionCount-arraylength  
  `regionCount` **must** be greater than `0`
- [](#VUID-VkCopyTensorInfoARM-commonparent)VUID-VkCopyTensorInfoARM-commonparent  
  Both of `dstTensor`, and `srcTensor` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html), [VkTensorCopyARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCopyARM.html), [vkCmdCopyTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyTensorARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCopyTensorInfoARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0