# VkExternalImageFormatProperties(3) Manual Page

## Name

VkExternalImageFormatProperties - Structure specifying supported external handle properties



## [](#_c_specification)C Specification

The `VkExternalImageFormatProperties` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkExternalImageFormatProperties {
    VkStructureType               sType;
    void*                         pNext;
    VkExternalMemoryProperties    externalMemoryProperties;
} VkExternalImageFormatProperties;
```

or the equivalent

```c++
// Provided by VK_KHR_external_memory_capabilities
typedef VkExternalImageFormatProperties VkExternalImageFormatPropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `externalMemoryProperties` is a [VkExternalMemoryProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryProperties.html) structure specifying various capabilities of the external handle type when used with the specified image creation parameters.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkExternalImageFormatProperties-sType-sType)VUID-VkExternalImageFormatProperties-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXTERNAL_IMAGE_FORMAT_PROPERTIES`

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkExternalMemoryProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryProperties.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExternalImageFormatProperties)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0