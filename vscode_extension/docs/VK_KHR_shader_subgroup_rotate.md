# VK\_KHR\_shader\_subgroup\_rotate(3) Manual Page

## Name

VK\_KHR\_shader\_subgroup\_rotate - device extension



## [](#_registered_extension_number)Registered Extension Number

417

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_subgroup\_rotate](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_subgroup_rotate.html)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.4](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.4-promotions)

## [](#_contact)Contact

- Kevin Petit [\[GitHub\]kpet](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shader_subgroup_rotate%5D%20%40kpet%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shader_subgroup_rotate%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_shader\_subgroup\_rotate](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_shader_subgroup_rotate.adoc)

**Last Modified Date**

2024-01-29

**IP Status**

No known IP claims.

**Contributors**

- Kévin Petit, Arm Ltd.
- Tobias Hector, AMD
- Jon Leech, Khronos
- Matthew Netsch, Qualcomm
- Jan-Harald Fredriksen, Arm Ltd.
- Graeme Leese, Broadcom
- Tom Olson, Arm Ltd.
- Spencer Fricke, LunarG Inc.

This extension adds support for the subgroup rotate instruction defined in SPV\_KHR\_subgroup\_rotate.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceShaderSubgroupRotateFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderSubgroupRotateFeaturesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_SHADER_SUBGROUP_ROTATE_EXTENSION_NAME`
- `VK_KHR_SHADER_SUBGROUP_ROTATE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SUBGROUP_ROTATE_FEATURES_KHR`
- Extending [VkSubgroupFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubgroupFeatureFlagBits.html):
  
  - `VK_SUBGROUP_FEATURE_ROTATE_BIT_KHR`
  - `VK_SUBGROUP_FEATURE_ROTATE_CLUSTERED_BIT_KHR`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [GroupNonUniformRotateKHR](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-GroupNonUniformRotateKHR)

## [](#_promotion_to_vulkan_1_4)Promotion to Vulkan 1.4

Functionality in this extension is included in core Vulkan 1.4 with the KHR suffix omitted. The original type, enum and command names are still available as aliases of the core functionality.

## [](#_version_history)Version History

- Revision 2, 2024-01-29 (Kévin Petit)
  
  - Add `VK_SUBGROUP_FEATURE_ROTATE_BIT_KHR` and `VK_SUBGROUP_FEATURE_ROTATE_CLUSTERED_BIT_KHR`
- Revision 1, 2023-06-20 (Kévin Petit)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_shader_subgroup_rotate)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0