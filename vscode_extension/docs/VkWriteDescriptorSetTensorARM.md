# VkWriteDescriptorSetTensorARM(3) Manual Page

## Name

VkWriteDescriptorSetTensorARM - Structure specifying descriptor tensor info



## [](#_c_specification)C Specification

The `VkWriteDescriptorSetTensorARM` structure is defined as:

```c++
// Provided by VK_ARM_tensors
typedef struct VkWriteDescriptorSetTensorARM {
    VkStructureType           sType;
    const void*               pNext;
    uint32_t                  tensorViewCount;
    const VkTensorViewARM*    pTensorViews;
} VkWriteDescriptorSetTensorARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `tensorViewCount` is the number of elements in `pTensorViews`.
- `pTensorViews` are the tensor views that will be used to update the descriptor set.

## [](#_description)Description

Valid Usage

- [](#VUID-VkWriteDescriptorSetTensorARM-nullDescriptor-09898)VUID-VkWriteDescriptorSetTensorARM-nullDescriptor-09898  
  If the [`nullDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nullDescriptor) feature is not enabled, each element of `pTensorViews` **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)

Valid Usage (Implicit)

- [](#VUID-VkWriteDescriptorSetTensorARM-sType-sType)VUID-VkWriteDescriptorSetTensorARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_TENSOR_ARM`
- [](#VUID-VkWriteDescriptorSetTensorARM-pTensorViews-parameter)VUID-VkWriteDescriptorSetTensorARM-pTensorViews-parameter  
  `pTensorViews` **must** be a valid pointer to an array of `tensorViewCount` valid [VkTensorViewARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewARM.html) handles
- [](#VUID-VkWriteDescriptorSetTensorARM-tensorViewCount-arraylength)VUID-VkWriteDescriptorSetTensorARM-tensorViewCount-arraylength  
  `tensorViewCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkTensorViewARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkWriteDescriptorSetTensorARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0