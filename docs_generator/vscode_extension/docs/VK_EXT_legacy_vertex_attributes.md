# VK\_EXT\_legacy\_vertex\_attributes(3) Manual Page

## Name

VK\_EXT\_legacy\_vertex\_attributes - device extension



## [](#_registered_extension_number)Registered Extension Number

496

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_EXT\_vertex\_input\_dynamic\_state](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_vertex_input_dynamic_state.html)

## [](#_special_use)Special Use

- [OpenGL / ES support](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Mike Blumenkrantz [\[GitHub\]zmike](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_legacy_vertex_attributes%5D%20%40zmike%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_legacy_vertex_attributes%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_legacy\_vertex\_attributes](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_legacy_vertex_attributes.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-02-23

**IP Status**

No known IP claims.

**Contributors**

- Mike Blumenkrantz, Valve
- Piers Daniell, NVIDIA
- Spencer Fricke, LunarG
- Alyssa Rosenzweig, Valve

## [](#_description)Description

This extension adds support for legacy features of (non-64-bit) vertex attributes as found in OpenGL:

- Vertex attributes loaded from arbitrary buffer alignments
- Vertex attributes using arbitrary strides
- Vertex attributes where the component data type of the binding does not match the component numeric type of the shader input

These features are only usable with dynamic vertex input. Unaligned loads of vertex attributes may incur performance penalties, indicated with a property.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceLegacyVertexAttributesFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLegacyVertexAttributesFeaturesEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceLegacyVertexAttributesPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLegacyVertexAttributesPropertiesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_LEGACY_VERTEX_ATTRIBUTES_EXTENSION_NAME`
- `VK_EXT_LEGACY_VERTEX_ATTRIBUTES_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LEGACY_VERTEX_ATTRIBUTES_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LEGACY_VERTEX_ATTRIBUTES_PROPERTIES_EXT`

## [](#_issues)Issues

1.) Should implementations convert float/integer values?

**RESOLVED**: No. When fetching an integer data type from float values or float data types from integer values, the resulting shader values are implementation-dependent.

## [](#_version_history)Version History

- Revision 1, 2024-02-16 (Mike Blumenkrantz)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_legacy_vertex_attributes)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0