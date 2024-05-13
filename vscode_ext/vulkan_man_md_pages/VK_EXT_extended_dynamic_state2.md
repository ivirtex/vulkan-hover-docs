# VK_EXT_extended_dynamic_state2(3) Manual Page

## Name

VK_EXT_extended_dynamic_state2 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

378

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.3-promotions"
  target="_blank" rel="noopener">Vulkan 1.3</a>

## <a href="#_contact" class="anchor"></a>Contact

- Vikram Kushwaha <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_extended_dynamic_state2%5D%20@vkushwaha-nv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_extended_dynamic_state2%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>vkushwaha-nv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-04-12

**IP Status**  
No known IP claims.

**Contributors**  
- Vikram Kushwaha, NVIDIA

- Piers Daniell, NVIDIA

- Jeff Bolz, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension adds some more dynamic state to support applications that
need to reduce the number of pipeline state objects they compile and
bind.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdSetDepthBiasEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthBiasEnableEXT.html)

- [vkCmdSetLogicOpEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLogicOpEXT.html)

- [vkCmdSetPatchControlPointsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetPatchControlPointsEXT.html)

- [vkCmdSetPrimitiveRestartEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetPrimitiveRestartEnableEXT.html)

- [vkCmdSetRasterizerDiscardEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnableEXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceExtendedDynamicState2FeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExtendedDynamicState2FeaturesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_EXTENDED_DYNAMIC_STATE_2_EXTENSION_NAME`

- `VK_EXT_EXTENDED_DYNAMIC_STATE_2_SPEC_VERSION`

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_DEPTH_BIAS_ENABLE_EXT`

  - `VK_DYNAMIC_STATE_LOGIC_OP_EXT`

  - `VK_DYNAMIC_STATE_PATCH_CONTROL_POINTS_EXT`

  - `VK_DYNAMIC_STATE_PRIMITIVE_RESTART_ENABLE_EXT`

  - `VK_DYNAMIC_STATE_RASTERIZER_DISCARD_ENABLE_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_DYNAMIC_STATE_2_FEATURES_EXT`

## <a href="#_promotion_to_vulkan_1_3" class="anchor"></a>Promotion to Vulkan 1.3

This extension has been partially promoted. The dynamic state enumerants
`VK_DYNAMIC_STATE_DEPTH_BIAS_ENABLE_EXT`,
`VK_DYNAMIC_STATE_PRIMITIVE_RESTART_ENABLE_EXT`, and
`VK_DYNAMIC_STATE_RASTERIZER_DISCARD_ENABLE_EXT`; and the corresponding
entry points in this extension are included in core Vulkan 1.3, with the
EXT suffix omitted. The enumerants and entry points for dynamic logic
operation and patch control points are not promoted, nor is the feature
structure. Extension interfaces that were promoted remain available as
aliases of the core functionality.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2021-04-12 (Vikram Kushwaha)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_extended_dynamic_state2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
