# VK\_NV\_fragment\_shading\_rate\_enums(3) Manual Page

## Name

VK\_NV\_fragment\_shading\_rate\_enums - device extension



## [](#_registered_extension_number)Registered Extension Number

327

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_fragment\_shading\_rate](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_fragment_shading_rate.html)

## [](#_contact)Contact

- Pat Brown [\[GitHub\]nvpbrown](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_fragment_shading_rate_enums%5D%20%40nvpbrown%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_fragment_shading_rate_enums%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2020-09-02

**Contributors**

- Pat Brown, NVIDIA
- Jeff Bolz, NVIDIA

## [](#_description)Description

This extension builds on the fragment shading rate functionality provided by the VK\_KHR\_fragment\_shading\_rate extension, adding support for “supersample” fragment shading rates that trigger multiple fragment shader invocations per pixel as well as a “no invocations” shading rate that discards any portions of a primitive that would use that shading rate.

## [](#_new_commands)New Commands

- [vkCmdSetFragmentShadingRateEnumNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetFragmentShadingRateEnumNV.html)

## [](#_new_structures)New Structures

- Extending [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html):
  
  - [VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceFragmentShadingRateEnumsFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFragmentShadingRateEnumsFeaturesNV.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV.html)

## [](#_new_enums)New Enums

- [VkFragmentShadingRateNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFragmentShadingRateNV.html)
- [VkFragmentShadingRateTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFragmentShadingRateTypeNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_FRAGMENT_SHADING_RATE_ENUMS_EXTENSION_NAME`
- `VK_NV_FRAGMENT_SHADING_RATE_ENUMS_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADING_RATE_ENUMS_FEATURES_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADING_RATE_ENUMS_PROPERTIES_NV`
  - `VK_STRUCTURE_TYPE_PIPELINE_FRAGMENT_SHADING_RATE_ENUM_STATE_CREATE_INFO_NV`

## [](#_issues)Issues

1. Why was this extension created? How should it be named?
   
   **RESOLVED**: The primary goal of this extension was to expose support for supersample and “no invocations” shading rates, which are supported by the VK\_NV\_shading\_rate\_image extension but not by VK\_KHR\_fragment\_shading\_rate. Because VK\_KHR\_fragment\_shading\_rate specifies the primitive shading rate using a fragment size in pixels, it lacks a good way to specify supersample rates. To deal with this, we defined enums covering shading rates supported by the KHR extension as well as the new shading rates and added structures and APIs accepting shading rate enums instead of fragment sizes.
   
   Since this extension adds two different types of shading rates, both expressed using enums, we chose the extension name VK\_NV\_fragment\_shading\_rate\_enums.
2. Is this a standalone extension?
   
   **RESOLVED**: No, this extension requires VK\_KHR\_fragment\_shading\_rate. In order to use the features of this extension, applications must enable the relevant features of KHR extension.
3. How are the shading rate enums used, and how were the enum values assigned?
   
   **RESOLVED**: The shading rates supported by the enums in this extension are accepted as pipeline, primitive, and attachment shading rates and behave identically. For the shading rates also supported by the KHR extension, the values assigned to the corresponding enums are identical to the values already used for the primitive and attachment shading rates in the KHR extension. For those enums, bits 0 and 1 specify the base two logarithm of the fragment height and bits 2 and 3 specify the base two logarithm of the fragment width. For the new shading rates added by this extension, we chose to use 11 through 14 (10 plus the base two logarithm of the invocation count) for the supersample rates and 15 for the “no invocations” rate. None of those values are supported as primitive or attachment shading rates by the KHR extension.
4. Between this extension, VK\_KHR\_fragment\_shading\_rate, and VK\_NV\_shading\_rate\_image, there are three different ways to specify shading rate state in a pipeline. How should we handle this?
   
   **RESOLVED**: We do not allow the concurrent use of VK\_NV\_shading\_rate\_image and VK\_KHR\_fragment\_shading\_rate; it is an error to enable shading rate features from both extensions. But we do allow applications to enable this extension together with VK\_KHR\_fragment\_shading\_rate together. While we expect that applications will never attach pipeline CreateInfo structures for both this extension and the KHR extension concurrently, Vulkan does not have any precedent forbidding such behavior and instead typically treats a pipeline created without an extension-specific CreateInfo structure as equivalent to one containing default values specified by the extension. Rather than adding such a rule considering the presence or absence of our new CreateInfo structure, we instead included a `shadingRateType` member to [VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html) that selects between using state specified by that structure and state specified by [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html).

## [](#_version_history)Version History

- Revision 1, 2020-09-02 (pbrown)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_fragment_shading_rate_enums)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0