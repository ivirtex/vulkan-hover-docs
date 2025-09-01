# VkExportMetalObjectCreateInfoEXT(3) Manual Page

## Name

VkExportMetalObjectCreateInfoEXT - Structure that identifies the Metal objects that can be exported from Vulkan objects



## [](#_c_specification)C Specification

To export Metal objects from Vulkan objects, the application **must** first indicate the intention to do so during the creation of the Vulkan object, by including one or more [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectCreateInfoEXT.html) structures in the `pNext` chain of the [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstanceCreateInfo.html), [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html), [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html), [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html), [VkBufferViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferViewCreateInfo.html), [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreCreateInfo.html), or [VkEventCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEventCreateInfo.html), in the corresponding Vulkan object creation command.

The `VkExportMetalObjectCreateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_metal_objects
typedef struct VkExportMetalObjectCreateInfoEXT {
    VkStructureType                       sType;
    const void*                           pNext;
    VkExportMetalObjectTypeFlagBitsEXT    exportObjectType;
} VkExportMetalObjectCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `exportObjectType` is a [VkExportMetalObjectTypeFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectTypeFlagBitsEXT.html) indicating the type of Metal object that the application may request to be exported from the Vulkan object.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkExportMetalObjectCreateInfoEXT-sType-sType)VUID-VkExportMetalObjectCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXPORT_METAL_OBJECT_CREATE_INFO_EXT`
- [](#VUID-VkExportMetalObjectCreateInfoEXT-exportObjectType-parameter)VUID-VkExportMetalObjectCreateInfoEXT-exportObjectType-parameter  
  If `exportObjectType` is not `0`, `exportObjectType` **must** be a valid [VkExportMetalObjectTypeFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectTypeFlagBitsEXT.html) value

## [](#_see_also)See Also

[VK\_EXT\_metal\_objects](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_metal_objects.html), [VkExportMetalObjectTypeFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectTypeFlagBitsEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExportMetalObjectCreateInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0