# VkExternalBufferProperties(3) Manual Page

## Name

VkExternalBufferProperties - Structure specifying supported external handle capabilities



## [](#_c_specification)C Specification

The `VkExternalBufferProperties` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkExternalBufferProperties {
    VkStructureType               sType;
    void*                         pNext;
    VkExternalMemoryProperties    externalMemoryProperties;
} VkExternalBufferProperties;
```

or the equivalent

```c++
// Provided by VK_KHR_external_memory_capabilities
typedef VkExternalBufferProperties VkExternalBufferPropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `externalMemoryProperties` is a [VkExternalMemoryProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryProperties.html) structure specifying various capabilities of the external handle type when used with the specified buffer creation parameters.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkExternalBufferProperties-sType-sType)VUID-VkExternalBufferProperties-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES`
- [](#VUID-VkExternalBufferProperties-pNext-pNext)VUID-VkExternalBufferProperties-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkExternalMemoryProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryProperties.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPhysicalDeviceExternalBufferProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceExternalBufferProperties.html), [vkGetPhysicalDeviceExternalBufferProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceExternalBufferProperties.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExternalBufferProperties).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0