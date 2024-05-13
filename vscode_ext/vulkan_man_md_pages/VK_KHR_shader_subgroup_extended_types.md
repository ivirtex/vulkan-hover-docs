# VK_KHR_shader_subgroup_extended_types(3) Manual Page

## Name

VK_KHR_shader_subgroup_extended_types - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

176

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[Version 1.1](#versions-1.1)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.2-promotions"
  target="_blank" rel="noopener">Vulkan 1.2</a>

## <a href="#_contact" class="anchor"></a>Contact

- Neil Henning <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shader_subgroup_extended_types%5D%20@sheredom%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shader_subgroup_extended_types%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>sheredom</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-01-08

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GLSL_EXT_shader_subgroup_extended_types`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_shader_subgroup_extended_types.txt)

**Contributors**  
- Jeff Bolz, NVIDIA

- Jan-Harald Fredriksen, Arm

- Neil Henning, AMD

- Daniel Koch, NVIDIA

- Jeff Leger, Qualcomm

- Graeme Leese, Broadcom

- David Neto, Google

- Daniel Rakos, AMD

## <a href="#_description" class="anchor"></a>Description

This extension enables the Non Uniform Group Operations in SPIR-V to
support 8-bit integer, 16-bit integer, 64-bit integer, 16-bit
floating-point, and vectors of these types.

## <a href="#_promotion_to_vulkan_1_2" class="anchor"></a>Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with
the KHR suffix omitted. The original type, enum and command names are
still available as aliases of the core functionality.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceShaderSubgroupExtendedTypesFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderSubgroupExtendedTypesFeaturesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_SHADER_SUBGROUP_EXTENDED_TYPES_EXTENSION_NAME`

- `VK_KHR_SHADER_SUBGROUP_EXTENDED_TYPES_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SUBGROUP_EXTENDED_TYPES_FEATURES_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-01-08 (Neil Henning)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_shader_subgroup_extended_types"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
