# VkTensorCopyARM(3) Manual Page

## Name

VkTensorCopyARM - Structure specifying an tensor copy region



## [](#_c_specification)C Specification

The `VkTensorCopyARM` structure is defined as:

```c++
// Provided by VK_ARM_tensors
typedef struct VkTensorCopyARM {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           dimensionCount;
    const uint64_t*    pSrcOffset;
    const uint64_t*    pDstOffset;
    const uint64_t*    pExtent;
} VkTensorCopyARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `dimensionCount` is the number of elements in the `pSrcOffset`, `pDstOffset` and `pExtent` arrays.
- `pSrcOffset` is `NULL` or an array of size `dimensionCount` providing an offset into the source tensor. When `pSrcOffset` is `NULL`, the offset into the source tensor is `0` in all dimensions.
- `pDstOffset` is `NULL` or an array of size `dimensionCount` providing an offset into the destination tensor. When `pDstOffset` is `NULL`, the offset into the destination tensor is `0` in all dimensions.
- `pExtent` is `NULL` or an array of size `dimensionCount` providing the number of elements to copy in each dimension. When `pExtent` is `NULL`, the number of elements to copy is taken as the total number of elements in each dimension of the source tensor.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkTensorCopyARM-sType-sType)VUID-VkTensorCopyARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_TENSOR_COPY_ARM`
- [](#VUID-VkTensorCopyARM-pNext-pNext)VUID-VkTensorCopyARM-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkTensorCopyARM-pSrcOffset-parameter)VUID-VkTensorCopyARM-pSrcOffset-parameter  
  If `pSrcOffset` is not `NULL`, `pSrcOffset` **must** be a valid pointer to an array of `dimensionCount` `uint64_t` values
- [](#VUID-VkTensorCopyARM-pDstOffset-parameter)VUID-VkTensorCopyARM-pDstOffset-parameter  
  If `pDstOffset` is not `NULL`, `pDstOffset` **must** be a valid pointer to an array of `dimensionCount` `uint64_t` values
- [](#VUID-VkTensorCopyARM-pExtent-parameter)VUID-VkTensorCopyARM-pExtent-parameter  
  If `pExtent` is not `NULL`, `pExtent` **must** be a valid pointer to an array of `dimensionCount` `uint64_t` values
- [](#VUID-VkTensorCopyARM-dimensionCount-arraylength)VUID-VkTensorCopyARM-dimensionCount-arraylength  
  `dimensionCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkCopyTensorInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyTensorInfoARM.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTensorCopyARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0