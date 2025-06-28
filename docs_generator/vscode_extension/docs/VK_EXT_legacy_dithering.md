# VK\_EXT\_legacy\_dithering(3) Manual Page

## Name

VK\_EXT\_legacy\_dithering - device extension



## [](#_registered_extension_number)Registered Extension Number

466

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_api_interactions)API Interactions

- Interacts with VK\_VERSION\_1\_3
- Interacts with VK\_VERSION\_1\_4
- Interacts with VK\_KHR\_dynamic\_rendering
- Interacts with VK\_KHR\_maintenance5

## [](#_special_use)Special Use

- [OpenGL / ES support](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Shahbaz Youssefi [\[GitHub\]syoussefi](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_legacy_dithering%5D%20%40syoussefi%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_legacy_dithering%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_legacy\_dithering](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_legacy_dithering.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-02-22

**Contributors**

- Shahbaz Youssefi, Google
- Graeme Leese, Broadcom
- Jan-Harald Fredriksen, Arm

## [](#_description)Description

This extension exposes a hardware feature used by some vendors to implement OpenGL’s dithering. The purpose of this extension is to support layering OpenGL over Vulkan, by allowing the layer to take advantage of the same hardware feature and provide equivalent dithering to OpenGL applications.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceLegacyDitheringFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLegacyDitheringFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_LEGACY_DITHERING_EXTENSION_NAME`
- `VK_EXT_LEGACY_DITHERING_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LEGACY_DITHERING_FEATURES_EXT`
- Extending [VkSubpassDescriptionFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescriptionFlagBits.html):
  
  - `VK_SUBPASS_DESCRIPTION_ENABLE_LEGACY_DITHERING_BIT_EXT`

If [VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html) or [Vulkan Version 1.3](#versions-1.3) and [VK\_KHR\_maintenance5](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance5.html) or [Vulkan Version 1.4](#versions-1.4) is supported:

- Extending [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html):
  
  - `VK_PIPELINE_CREATE_2_ENABLE_LEGACY_DITHERING_BIT_EXT`
- Extending [VkRenderingFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingFlagBits.html):
  
  - `VK_RENDERING_ENABLE_LEGACY_DITHERING_BIT_EXT`

## [](#_version_history)Version History

- Revision 1, 2022-03-31 (Shahbaz Youssefi)
  
  - Internal revisions
- Revision 2, 2024-02-22 (Shahbaz Youssefi)
  
  - Added pipeline create flag to support dynamic rendering

## [](#_issues)Issues

1\) In OpenGL, the dither state can change dynamically. Should this extension add a pipeline state for dither?

**RESOLVED**: No. Changing dither state is rarely, if ever, done during rendering. Every surveyed Android application either entirely disables dither, explicitly enables it, or uses the default state (which is enabled). Additionally, on some hardware dither can only be specified in a render pass granularity, so a change in dither state would necessarily need to cause a render pass break. This extension considers dynamic changes in OpenGL dither state a theoretical situation, and expects the layer to break the render pass in such a situation without any practical downsides.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_legacy_dithering)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0