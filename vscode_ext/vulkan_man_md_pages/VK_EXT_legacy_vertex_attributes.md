# VK_EXT_legacy_vertex_attributes(3) Manual Page

## Name

VK_EXT_legacy_vertex_attributes - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

496

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_EXT_vertex_input_dynamic_state](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_vertex_input_dynamic_state.html)  

## <a href="#_special_use" class="anchor"></a>Special Use

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-compatibility-specialuse"
  target="_blank" rel="noopener">OpenGL / ES support</a>

## <a href="#_contact" class="anchor"></a>Contact

- Mike Blumenkrantz <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_legacy_vertex_attributes%5D%20@zmike%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_legacy_vertex_attributes%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>zmike</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_legacy_vertex_attributes](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_legacy_vertex_attributes.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2024-02-23

**IP Status**  
No known IP claims.

**Contributors**  
- Mike Blumenkrantz, Valve

- Piers Daniell, NVIDIA

- Spencer Fricke, LunarG

- Alyssa Rosenzweig, Valve

## <a href="#_description" class="anchor"></a>Description

This extension adds support for legacy features of (non-64-bit) vertex
attributes as found in OpenGL:

- Vertex attributes loaded from arbitrary buffer alignments

- Vertex attributes using arbitrary strides

- Vertex attributes where the component data type of the binding does
  not match the component numeric type of the shader input

These features are only usable with dynamic vertex input. Unaligned
loads of vertex attributes may incur performance penalties, indicated
with a property.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceLegacyVertexAttributesFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLegacyVertexAttributesFeaturesEXT.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceLegacyVertexAttributesPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLegacyVertexAttributesPropertiesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_LEGACY_VERTEX_ATTRIBUTES_EXTENSION_NAME`

- `VK_EXT_LEGACY_VERTEX_ATTRIBUTES_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LEGACY_VERTEX_ATTRIBUTES_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LEGACY_VERTEX_ATTRIBUTES_PROPERTIES_EXT`

## <a href="#_issues" class="anchor"></a>Issues

1.) Should implementations convert float/integer values?

**RESOLVED**: No. When fetching an integer data type from float values
or float data types from integer values, the resulting shader values are
implementation-dependent.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2024-02-16 (Mike Blumenkrantz)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_legacy_vertex_attributes"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
