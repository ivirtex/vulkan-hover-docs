# vkExportMetalObjectsEXT(3) Manual Page

## Name

vkExportMetalObjectsEXT - Export Metal objects from the corresponding
Vulkan objects



## <a href="#_c_specification" class="anchor"></a>C Specification

To export Metal objects that underlie Vulkan objects, call:

``` c
// Provided by VK_EXT_metal_objects
void vkExportMetalObjectsEXT(
    VkDevice                                    device,
    VkExportMetalObjectsInfoEXT*                pMetalObjectsInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device that created the Vulkan objects.

- `pMetalObjectsInfo` is a pointer to a
  [VkExportMetalObjectsInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectsInfoEXT.html)
  structure whose `pNext` chain contains structures, each identifying a
  Vulkan object and providing a pointer through which the Metal object
  will be returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkExportMetalObjectsEXT-device-parameter"
  id="VUID-vkExportMetalObjectsEXT-device-parameter"></a>
  VUID-vkExportMetalObjectsEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkExportMetalObjectsEXT-pMetalObjectsInfo-parameter"
  id="VUID-vkExportMetalObjectsEXT-pMetalObjectsInfo-parameter"></a>
  VUID-vkExportMetalObjectsEXT-pMetalObjectsInfo-parameter  
  `pMetalObjectsInfo` **must** be a valid pointer to a
  [VkExportMetalObjectsInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectsInfoEXT.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_metal_objects](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_metal_objects.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkExportMetalObjectsInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectsInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkExportMetalObjectsEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
