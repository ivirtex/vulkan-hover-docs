# VkExportMetalObjectCreateInfoEXT(3) Manual Page

## Name

VkExportMetalObjectCreateInfoEXT - Structure that identifies the Metal
objects that can be exported from Vulkan objects



## <a href="#_c_specification" class="anchor"></a>C Specification

To export Metal objects from Vulkan objects, the application **must**
first indicate the intention to do so during the creation of the Vulkan
object, by including one or more
[VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html)
structures in the `pNext` chain of the
[VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateInfo.html),
[VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html),
[VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html),
[VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html),
[VkBufferViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferViewCreateInfo.html),
[VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreCreateInfo.html), or
[VkEventCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEventCreateInfo.html), in the corresponding Vulkan
object creation command.

The `VkExportMetalObjectCreateInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_metal_objects
typedef struct VkExportMetalObjectCreateInfoEXT {
    VkStructureType                       sType;
    const void*                           pNext;
    VkExportMetalObjectTypeFlagBitsEXT    exportObjectType;
} VkExportMetalObjectCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `exportObjectType` is a
  [VkExportMetalObjectTypeFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectTypeFlagBitsEXT.html)
  indicating the type of Metal object that the application may request
  to be exported from the Vulkan object.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkExportMetalObjectCreateInfoEXT-sType-sType"
  id="VUID-VkExportMetalObjectCreateInfoEXT-sType-sType"></a>
  VUID-VkExportMetalObjectCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_EXPORT_METAL_OBJECT_CREATE_INFO_EXT`

- <a
  href="#VUID-VkExportMetalObjectCreateInfoEXT-exportObjectType-parameter"
  id="VUID-VkExportMetalObjectCreateInfoEXT-exportObjectType-parameter"></a>
  VUID-VkExportMetalObjectCreateInfoEXT-exportObjectType-parameter  
  If `exportObjectType` is not `0`, `exportObjectType` **must** be a
  valid
  [VkExportMetalObjectTypeFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectTypeFlagBitsEXT.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_metal_objects](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_metal_objects.html),
[VkExportMetalObjectTypeFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectTypeFlagBitsEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExportMetalObjectCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
