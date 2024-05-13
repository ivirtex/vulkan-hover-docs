# VK_EXT_legacy_dithering(3) Manual Page

## Name

VK_EXT_legacy_dithering - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

466

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_VERSION_1_3

- Interacts with VK_KHR_dynamic_rendering

- Interacts with VK_KHR_maintenance5

## <a href="#_special_use" class="anchor"></a>Special Use

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-compatibility-specialuse"
  target="_blank" rel="noopener">OpenGL / ES support</a>

## <a href="#_contact" class="anchor"></a>Contact

- Shahbaz Youssefi <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_legacy_dithering%5D%20@syoussefi%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_legacy_dithering%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>syoussefi</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_legacy_dithering](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_legacy_dithering.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2024-02-22

**Contributors**  
- Shahbaz Youssefi, Google

- Graeme Leese, Broadcom

- Jan-Harald Fredriksen, Arm

## <a href="#_description" class="anchor"></a>Description

This extension exposes a hardware feature used by some vendors to
implement OpenGL’s dithering. The purpose of this extension is to
support layering OpenGL over Vulkan, by allowing the layer to take
advantage of the same hardware feature and provide equivalent dithering
to OpenGL applications.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceLegacyDitheringFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLegacyDitheringFeaturesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_LEGACY_DITHERING_EXTENSION_NAME`

- `VK_EXT_LEGACY_DITHERING_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LEGACY_DITHERING_FEATURES_EXT`

- Extending
  [VkSubpassDescriptionFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescriptionFlagBits.html):

  - `VK_SUBPASS_DESCRIPTION_ENABLE_LEGACY_DITHERING_BIT_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2022-03-31 (Shahbaz Youssefi)

  - Internal revisions

- Revision 2, 2024-02-22 (Shahbaz Youssefi)

  - Added pipeline create flag to support dynamic rendering

## <a href="#_issues" class="anchor"></a>Issues

1\) In OpenGL, the dither state can change dynamically. Should this
extension add a pipeline state for dither?

**RESOLVED**: No. Changing dither state is rarely, if ever, done during
rendering. Every surveyed Android application either entirely disables
dither, explicitly enables it, or uses the default state (which is
enabled). Additionally, on some hardware dither can only be specified in
a render pass granularity, so a change in dither state would necessarily
need to cause a render pass break. This extension considers dynamic
changes in OpenGL dither state a theoretical situation, and expects the
layer to break the render pass in such a situation without any practical
downsides.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_legacy_dithering"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
