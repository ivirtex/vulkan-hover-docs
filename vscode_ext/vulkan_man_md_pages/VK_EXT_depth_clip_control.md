# VK_EXT_depth_clip_control(3) Manual Page

## Name

VK_EXT_depth_clip_control - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

356

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_special_use" class="anchor"></a>Special Use

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-compatibility-specialuse"
  target="_blank" rel="noopener">OpenGL / ES support</a>

## <a href="#_contact" class="anchor"></a>Contact

- Shahbaz Youssefi <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_depth_clip_control%5D%20@syoussefi%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_depth_clip_control%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>syoussefi</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-11-09

**Contributors**  
- Spencer Fricke, Samsung Electronics

- Shahbaz Youssefi, Google

- Ralph Potter, Samsung Electronics

## <a href="#_description" class="anchor"></a>Description

This extension allows the application to use the OpenGL depth range in
NDC, i.e. with depth in range \[-1, 1\], as opposed to Vulkan’s default
of \[0, 1\]. The purpose of this extension is to allow efficient
layering of OpenGL over Vulkan, by avoiding emulation in the
pre-rasterization shader stages. This emulation, which effectively
duplicates gl_Position but with a different depth value, costs ALU and
consumes shader output components that the implementation may not have
to spare to meet OpenGL minimum requirements.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceDepthClipControlFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDepthClipControlFeaturesEXT.html)

- Extending
  [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html):

  - [VkPipelineViewportDepthClipControlCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportDepthClipControlCreateInfoEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_DEPTH_CLIP_CONTROL_EXTENSION_NAME`

- `VK_EXT_DEPTH_CLIP_CONTROL_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_CLIP_CONTROL_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_DEPTH_CLIP_CONTROL_CREATE_INFO_EXT`

## <a href="#_issues" class="anchor"></a>Issues

1\) Should this extension include an origin control option to match
GL_LOWER_LEFT found in ARB_clip_control?

**RESOLVED**: No. The fix for porting over the origin is a simple y-axis
flip. The depth clip control is a much harder problem to solve than what
this extension is aimed to solve. Adding an equivalent to GL_LOWER_LEFT
would require more testing.

2\) Should this pipeline state be dynamic?

**RESOLVED**: Yes. The purpose of this extension is to emulate the
OpenGL depth range, which is expected to be globally fixed (in case of
OpenGL ES) or very infrequently changed (with `glClipControl` in
OpenGL).

3\) Should the control provided in this extension be an enum that could
be extended in the future?

**RESOLVED**: No. It is highly unlikely that the depth range is changed
to anything other than \[0, 1\] in the future. Should that happen a new
extension will be required to extend such an enum, and that extension
might as well add a new struct to chain to
[VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html)::`pNext`
instead.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 0, 2020-10-01 (Spencer Fricke)

  - Internal revisions

- Revision 1, 2020-11-26 (Shahbaz Youssefi)

  - Language fixes

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_depth_clip_control"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
