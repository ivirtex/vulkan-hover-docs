# VK\_EXT\_image\_robustness(3) Manual Page

## Name

VK\_EXT\_image\_robustness - device extension



## [](#_registered_extension_number)Registered Extension Number

336

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.3](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.3-promotions)

## [](#_contact)Contact

- Graeme Leese [\[GitHub\]gnl21](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_image_robustness%5D%20%40gnl21%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_image_robustness%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2020-04-27

**IP Status**

No known IP claims.

**Contributors**

- Graeme Leese, Broadcom
- Jan-Harald Fredriksen, ARM
- Jeff Bolz, NVIDIA
- Spencer Fricke, Samsung
- Courtney Goeltzenleuchter, Google
- Slawomir Cygan, Intel

## [](#_description)Description

This extension adds stricter requirements for how out of bounds reads from images are handled. Rather than returning undefined values, most out of bounds reads return R, G, and B values of zero and alpha values of either zero or one. Components not present in the image format may be set to zero or to values based on the format as described in [Conversion to RGBA](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-conversion-to-rgba).

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceImageRobustnessFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageRobustnessFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_IMAGE_ROBUSTNESS_EXTENSION_NAME`
- `VK_EXT_IMAGE_ROBUSTNESS_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_ROBUSTNESS_FEATURES_EXT`

## [](#_promotion_to_vulkan_1_3)Promotion to Vulkan 1.3

Vulkan APIs in this extension are included in core Vulkan 1.3, with the EXT suffix omitted. External interactions defined by this extension, such as SPIR-V token names, retain their original names. The original Vulkan API names are still available as aliases of the core functionality.

## [](#_issues)Issues

1. How does this extension differ from VK\_EXT\_robustness2?

The guarantees provided by this extension are a subset of those provided by the robustImageAccess2 feature of VK\_EXT\_robustness2. Where this extension allows return values of (0, 0, 0, 0) or (0, 0, 0, 1), robustImageAccess2 requires that a particular value dependent on the image format be returned. This extension provides no guarantees about the values returned for an access to an invalid Lod.

## [](#_examples)Examples

None.

## [](#_version_history)Version History

- Revision 1, 2020-04-27 (Graeme Leese)
- Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_image_robustness)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0