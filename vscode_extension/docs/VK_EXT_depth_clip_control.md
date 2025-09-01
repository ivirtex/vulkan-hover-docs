# VK\_EXT\_depth\_clip\_control(3) Manual Page

## Name

VK\_EXT\_depth\_clip\_control - device extension



## [](#_registered_extension_number)Registered Extension Number

356

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_special_use)Special Use

- [OpenGL / ES support](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Shahbaz Youssefi [\[GitHub\]syoussefi](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_depth_clip_control%5D%20%40syoussefi%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_depth_clip_control%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-11-09

**Contributors**

- Spencer Fricke, Samsung Electronics
- Shahbaz Youssefi, Google
- Ralph Potter, Samsung Electronics

## [](#_description)Description

This extension allows the application to use the OpenGL depth range in NDC, i.e. with depth in range \[-1, 1], as opposed to Vulkan’s default of \[0, 1]. The purpose of this extension is to allow efficient layering of OpenGL over Vulkan, by avoiding emulation in the pre-rasterization shader stages. This emulation, which effectively duplicates gl\_Position but with a different depth value, costs ALU and consumes shader output components that the implementation may not have to spare to meet OpenGL minimum requirements.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceDepthClipControlFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDepthClipControlFeaturesEXT.html)
- Extending [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportStateCreateInfo.html):
  
  - [VkPipelineViewportDepthClipControlCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportDepthClipControlCreateInfoEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_DEPTH_CLIP_CONTROL_EXTENSION_NAME`
- `VK_EXT_DEPTH_CLIP_CONTROL_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_CLIP_CONTROL_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_DEPTH_CLIP_CONTROL_CREATE_INFO_EXT`

## [](#_issues)Issues

1\) Should this extension include an origin control option to match GL\_LOWER\_LEFT found in ARB\_clip\_control?

**RESOLVED**: No. The fix for porting over the origin is a simple y-axis flip. The depth clip control is a much harder problem to solve than what this extension is aimed to solve. Adding an equivalent to GL\_LOWER\_LEFT would require more testing.

2\) Should this pipeline state be dynamic?

**RESOLVED**: Yes. The purpose of this extension is to emulate the OpenGL depth range, which is expected to be globally fixed (in case of OpenGL ES) or very infrequently changed (with `glClipControl` in OpenGL).

3\) Should the control provided in this extension be an enum that could be extended in the future?

**RESOLVED**: No. It is highly unlikely that the depth range is changed to anything other than \[0, 1] in the future. Should that happen a new extension will be required to extend such an enum, and that extension might as well add a new structure to chain to [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportStateCreateInfo.html)::`pNext` instead.

## [](#_version_history)Version History

- Revision 0, 2020-10-01 (Spencer Fricke)
  
  - Internal revisions
- Revision 1, 2020-11-26 (Shahbaz Youssefi)
  
  - Language fixes

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_depth_clip_control).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0