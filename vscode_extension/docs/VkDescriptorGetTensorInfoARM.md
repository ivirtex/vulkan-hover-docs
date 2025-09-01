# VkDescriptorGetTensorInfoARM(3) Manual Page

## Name

VkDescriptorGetTensorInfoARM - Structure specifying parameters to get descriptor data for tensor views



## [](#_c_specification)C Specification

The `VkDescriptorGetTensorInfoARM` is defined as:

```c++
// Provided by VK_EXT_descriptor_buffer with VK_ARM_tensors
typedef struct VkDescriptorGetTensorInfoARM {
    VkStructureType    sType;
    const void*        pNext;
    VkTensorViewARM    tensorView;
} VkDescriptorGetTensorInfoARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `tensorView` is a [VkTensorViewARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewARM.html) handle specifying the parameters of a `VK_DESCRIPTOR_TYPE_TENSOR_ARM` descriptor.

## [](#_description)Description

Valid Usage

- [](#VUID-VkDescriptorGetTensorInfoARM-nullDescriptor-09899)VUID-VkDescriptorGetTensorInfoARM-nullDescriptor-09899  
  If the [`nullDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nullDescriptor) feature is not enabled, `tensorView` **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)

Valid Usage (Implicit)

- [](#VUID-VkDescriptorGetTensorInfoARM-sType-sType)VUID-VkDescriptorGetTensorInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DESCRIPTOR_GET_TENSOR_INFO_ARM`
- [](#VUID-VkDescriptorGetTensorInfoARM-tensorView-parameter)VUID-VkDescriptorGetTensorInfoARM-tensorView-parameter  
  `tensorView` **must** be a valid [VkTensorViewARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewARM.html) handle

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkTensorViewARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDescriptorGetTensorInfoARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0