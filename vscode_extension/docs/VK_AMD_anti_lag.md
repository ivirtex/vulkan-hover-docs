# VK\_AMD\_anti\_lag(3) Manual Page

## Name

VK\_AMD\_anti\_lag - device extension



## [](#_registered_extension_number)Registered Extension Number

477

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_contact)Contact

- Stu Smith

## [](#_extension_proposal)Extension Proposal

[VK\_AMD\_anti\_lag](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_AMD_anti_lag.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-06-06

**IP Status**

No known IP claims.

**Contributors**

- Tobias Hector, AMD
- Stuart Smith, AMD
- Arkadiusz Sarwa, AMD

## [](#_description)Description

This extension automatically paces the CPU to make sure it does not get too far ahead of the GPU, reducing the latency between inputs received and updates on the screen. Additionally, Anti-Lag+ offers applications the ability to inform the driver when input processing begins, in order to align the timing of display updates, enabling even lower latency between receiving input and displaying on the screen.

## [](#_new_commands)New Commands

- [vkAntiLagUpdateAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAntiLagUpdateAMD.html)

## [](#_new_structures)New Structures

- [VkAntiLagDataAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAntiLagDataAMD.html)
- [VkAntiLagPresentationInfoAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAntiLagPresentationInfoAMD.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceAntiLagFeaturesAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceAntiLagFeaturesAMD.html)

## [](#_new_enums)New Enums

- [VkAntiLagModeAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAntiLagModeAMD.html)
- [VkAntiLagStageAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAntiLagStageAMD.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_AMD_ANTI_LAG_EXTENSION_NAME`
- `VK_AMD_ANTI_LAG_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_ANTI_LAG_DATA_AMD`
  - `VK_STRUCTURE_TYPE_ANTI_LAG_PRESENTATION_INFO_AMD`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ANTI_LAG_FEATURES_AMD`

## [](#_version_history)Version History

- Revision 1, 2024-06-06 (Arkadiusz Sarw)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_AMD_anti_lag).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0