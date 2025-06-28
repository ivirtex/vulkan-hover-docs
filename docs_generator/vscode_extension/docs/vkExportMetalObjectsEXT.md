# vkExportMetalObjectsEXT(3) Manual Page

## Name

vkExportMetalObjectsEXT - Export Metal objects from the corresponding Vulkan objects



## [](#_c_specification)C Specification

To export Metal objects that underlie Vulkan objects, call:

```c++
// Provided by VK_EXT_metal_objects
void vkExportMetalObjectsEXT(
    VkDevice                                    device,
    VkExportMetalObjectsInfoEXT*                pMetalObjectsInfo);
```

## [](#_parameters)Parameters

- `device` is the device that created the Vulkan objects.
- `pMetalObjectsInfo` is a pointer to a [VkExportMetalObjectsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectsInfoEXT.html) structure whose `pNext` chain contains structures, each identifying a Vulkan object and providing a pointer through which the Metal object will be returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkExportMetalObjectsEXT-device-parameter)VUID-vkExportMetalObjectsEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkExportMetalObjectsEXT-pMetalObjectsInfo-parameter)VUID-vkExportMetalObjectsEXT-pMetalObjectsInfo-parameter  
  `pMetalObjectsInfo` **must** be a valid pointer to a [VkExportMetalObjectsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectsInfoEXT.html) structure

## [](#_see_also)See Also

[VK\_EXT\_metal\_objects](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_metal_objects.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkExportMetalObjectsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectsInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkExportMetalObjectsEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0