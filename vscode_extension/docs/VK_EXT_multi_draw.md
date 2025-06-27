# VK_EXT_multi_draw(3) Manual Page

## Name

VK_EXT_multi_draw - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

393

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Mike Blumenkrantz <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_multi_draw%5D%20@zmike%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_multi_draw%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>zmike</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-05-19

**Interactions and External Dependencies**  
- Interacts with Vulkan 1.1.

- Interacts with
  [`VK_KHR_shader_draw_parameters`](VK_KHR_shader_draw_parameters.html).

**IP Status**  
No known IP claims.

**Contributors**  
- Mike Blumenkrantz, VALVE

- Piers Daniell, NVIDIA

- Faith Ekstrand, INTEL

- Spencer Fricke, SAMSUNG

- Ricardo Garcia, IGALIA

- Jon Leech, KHRONOS

- Stu Smith, AMD

## <a href="#_description" class="anchor"></a>Description

Processing multiple draw commands in sequence incurs measurable overhead
within drivers due to repeated state checks and updates during dispatch.
This extension enables passing the entire sequence of draws directly to
the driver in order to avoid any such overhead, using an array of
[VkMultiDrawInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultiDrawInfoEXT.html) or
[VkMultiDrawIndexedInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultiDrawIndexedInfoEXT.html) structs with
`vkCmdDrawMultiEXT` or `vkCmdDrawMultiIndexedEXT`, respectively. These
functions could be used any time multiple draw commands are being
recorded without any state changes between them in order to maximize
performance.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdDrawMultiEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawMultiEXT.html)

- [vkCmdDrawMultiIndexedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawMultiIndexedEXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkMultiDrawIndexedInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultiDrawIndexedInfoEXT.html)

- [VkMultiDrawInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultiDrawInfoEXT.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceMultiDrawFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMultiDrawFeaturesEXT.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceMultiDrawPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMultiDrawPropertiesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_MULTI_DRAW_EXTENSION_NAME`

- `VK_EXT_MULTI_DRAW_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTI_DRAW_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTI_DRAW_PROPERTIES_EXT`

## <a href="#_new_or_modified_built_in_variables" class="anchor"></a>New or Modified Built-In Variables

- (modified)`DrawIndex`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2021-01-20 (Mike Blumenkrantz)

  - Initial version

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_multi_draw"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
