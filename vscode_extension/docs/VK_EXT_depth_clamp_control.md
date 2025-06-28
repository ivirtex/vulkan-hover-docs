# VK\_EXT\_depth\_clamp\_control(3) Manual Page

## Name

VK\_EXT\_depth\_clamp\_control - device extension



## [](#_registered_extension_number)Registered Extension Number

583

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Jules Blok [\[GitHub\]jules](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_depth_clamp_control%5D%20%40jules%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_depth_clamp_control%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_depth\_clamp\_control](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_depth_clamp_control.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-07-15

**Contributors**

- Jules Blok, Independent

## [](#_description)Description

This extension allows the application to control the viewport depth clamp range separately from the viewport `minDepth` and `maxDepth`. This gives the ability for the application to restrict depth values to an application-defined range rather than the viewport depth range or the range defined in the [VK\_EXT\_depth\_clamp\_zero\_one](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_depth_clamp_zero_one.html) extension.

It can be used to set a smaller or larger clamping range than the viewport depth range without affecting the depth mapping of the viewport transform. Another possible use of this extension is to restrict depth values beyond the viewport depth range to a clamping range other than the \[0, 1] range defined in the [VK\_EXT\_depth\_clamp\_zero\_one](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_depth_clamp_zero_one.html) extension.

## [](#_new_commands)New Commands

- [vkCmdSetDepthClampRangeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthClampRangeEXT.html)

## [](#_new_structures)New Structures

- [VkDepthClampRangeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDepthClampRangeEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceDepthClampControlFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDepthClampControlFeaturesEXT.html)
- Extending [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportStateCreateInfo.html):
  
  - [VkPipelineViewportDepthClampControlCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportDepthClampControlCreateInfoEXT.html)

## [](#_new_enums)New Enums

- [VkDepthClampModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDepthClampModeEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_DEPTH_CLAMP_CONTROL_EXTENSION_NAME`
- `VK_EXT_DEPTH_CLAMP_CONTROL_SPEC_VERSION`
- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_DEPTH_CLAMP_RANGE_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_CLAMP_CONTROL_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_DEPTH_CLAMP_CONTROL_CREATE_INFO_EXT`

## [](#_issues)Issues

1\) Should the depth clamp range be a per-viewport parameter?

**RESOLVED**: No. Because the depth clamp range was previously defined to be equal to the viewport depth range, conformant runtimes are already handling the depth clamp range as a per-viewport parameter. However because of complexities from interactions with multiple viewports a per-viewport clamp range is left to a future extensions if a use case arises.

2\) Should this pipeline state be dynamic?

**RESOLVED**: Yes. Since the viewport depth range can already be a dynamic state conformant runtimes are already able to handle the depth clamp range as a dynamic state.

3\) Can the depth clamp range be ignored when depth clamping is disabled?

**RESOLVED**: Yes. This extension overrides the clamping range used only when depth clamping is enabled. The alternative would be highly unintuitive. As a consequence the [VK\_EXT\_depth\_clip\_enable](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_depth_clip_enable.html) extension is required if depth clipping is desired in combination with this extension.

## [](#_version_history)Version History

- Revision 1, 2024-02-13 (Jules Blok)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_depth_clamp_control)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0