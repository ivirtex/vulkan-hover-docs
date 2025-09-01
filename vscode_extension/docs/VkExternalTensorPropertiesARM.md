# VkExternalTensorPropertiesARM(3) Manual Page

## Name

VkExternalTensorPropertiesARM - Structure specifying supported external handle capabilities for a tensor



## [](#_c_specification)C Specification

The `VkExternalTensorPropertiesARM` structure is defined as:

```c++
// Provided by VK_ARM_tensors
typedef struct VkExternalTensorPropertiesARM {
    VkStructureType               sType;
    const void*                   pNext;
    VkExternalMemoryProperties    externalMemoryProperties;
} VkExternalTensorPropertiesARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `externalMemoryProperties` is a [VkExternalMemoryProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryProperties.html) structure specifying various capabilities of the external handle type when used with the specified tensor creation parameters.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkExternalTensorPropertiesARM-sType-sType)VUID-VkExternalTensorPropertiesARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXTERNAL_TENSOR_PROPERTIES_ARM`
- [](#VUID-VkExternalTensorPropertiesARM-pNext-pNext)VUID-VkExternalTensorPropertiesARM-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkExternalTensorPropertiesARM-externalMemoryProperties-parameter)VUID-VkExternalTensorPropertiesARM-externalMemoryProperties-parameter  
  `externalMemoryProperties` **must** be a valid [VkExternalMemoryProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryProperties.html) structure

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkExternalMemoryProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryProperties.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPhysicalDeviceExternalTensorPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceExternalTensorPropertiesARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExternalTensorPropertiesARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0