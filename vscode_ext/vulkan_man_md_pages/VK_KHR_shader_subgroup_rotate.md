# VK_KHR_shader_subgroup_rotate(3) Manual Page

## Name

VK_KHR_shader_subgroup_rotate - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

417

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_KHR_subgroup_rotate](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_subgroup_rotate.html)

## <a href="#_contact" class="anchor"></a>Contact

- Kevin Petit <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shader_subgroup_rotate%5D%20@kpet%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shader_subgroup_rotate%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>kpet</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_KHR_shader_subgroup_rotate](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_shader_subgroup_rotate.adoc)

**Last Modified Date**  
2024-01-29

**IP Status**  
No known IP claims.

**Contributors**  
- Kévin Petit, Arm Ltd.

- Tobias Hector, AMD

- John Leech, Khronos

- Matthew Netsch, Qualcomm

- Jan-Harald Fredriksen, Arm Ltd.

- Graeme Leese, Broadcom

- Tom Olson, Arm Ltd.

- Spencer Fricke, LunarG Inc.

This extension adds support for the subgroup rotate instruction defined
in SPV_KHR_subgroup_rotate.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceShaderSubgroupRotateFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderSubgroupRotateFeaturesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_SHADER_SUBGROUP_ROTATE_EXTENSION_NAME`

- `VK_KHR_SHADER_SUBGROUP_ROTATE_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SUBGROUP_ROTATE_FEATURES_KHR`

- Extending [VkSubgroupFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubgroupFeatureFlagBits.html):

  - `VK_SUBGROUP_FEATURE_ROTATE_BIT_KHR`

  - `VK_SUBGROUP_FEATURE_ROTATE_CLUSTERED_BIT_KHR`

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-GroupNonUniformRotateKHR"
  target="_blank" rel="noopener">GroupNonUniformRotateKHR</a>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 2, 2024-01-29 (Kévin Petit)

  - Add `VK_SUBGROUP_FEATURE_ROTATE_BIT_KHR` and
    `VK_SUBGROUP_FEATURE_ROTATE_CLUSTERED_BIT_KHR`

- Revision 1, 2023-06-20 (Kévin Petit)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_shader_subgroup_rotate"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
