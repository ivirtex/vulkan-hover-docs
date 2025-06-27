# VK_NVX_multiview_per_view_attributes(3) Manual Page

## Name

VK_NVX_multiview_per_view_attributes - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

98

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_multiview](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_multiview.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_NVX_multiview_per_view_attributes](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/NVX/SPV_NVX_multiview_per_view_attributes.html)

## <a href="#_contact" class="anchor"></a>Contact

- Jeff Bolz <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NVX_multiview_per_view_attributes%5D%20@jeffbolznv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NVX_multiview_per_view_attributes%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>jeffbolznv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-01-13

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_NVX_multiview_per_view_attributes`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nvx/GL_NVX_multiview_per_view_attributes.txt)

- This extension interacts with
  [`VK_NV_viewport_array2`](VK_NV_viewport_array2.html).

**Contributors**  
- Jeff Bolz, NVIDIA

- Daniel Koch, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension adds a new way to write shaders to be used with multiview
subpasses, where the attributes for all views are written out by a
single invocation of the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization shader stages</a>.
Related SPIR-V and GLSL extensions
`SPV_NVX_multiview_per_view_attributes` and
`GL_NVX_multiview_per_view_attributes` introduce per-view position and
viewport mask attributes arrays, and this extension defines how those
per-view attribute arrays are interpreted by Vulkan. Pipelines using
per-view attributes **may** only execute the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization shader stages</a> once
for all views rather than once per-view, which reduces redundant shading
work.

A subpass creation flag controls whether the subpass uses this
extension. A subpass **must** either exclusively use this extension or
not use it at all.

Some Vulkan implementations only support the position attribute varying
between views in the X component. A subpass can declare via a second
creation flag whether all pipelines compiled for this subpass will obey
this restriction.

Shaders that use the new per-view outputs (e.g. `gl_PositionPerViewNV`)
**must** also write the non-per-view output (`gl_Position`), and the
values written **must** be such that
`gl_Position = gl_PositionPerViewNV[gl_ViewIndex]` for all views in the
subpass. Implementations are free to either use the per-view outputs or
the non-per-view outputs, whichever would be more efficient.

If [`VK_NV_viewport_array2`](VK_NV_viewport_array2.html) is not also
supported and enabled, the per-view viewport mask **must** not be used.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NVX_MULTIVIEW_PER_VIEW_ATTRIBUTES_EXTENSION_NAME`

- `VK_NVX_MULTIVIEW_PER_VIEW_ATTRIBUTES_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PER_VIEW_ATTRIBUTES_PROPERTIES_NVX`

- Extending
  [VkSubpassDescriptionFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescriptionFlagBits.html):

  - `VK_SUBPASS_DESCRIPTION_PER_VIEW_ATTRIBUTES_BIT_NVX`

  - `VK_SUBPASS_DESCRIPTION_PER_VIEW_POSITION_X_ONLY_BIT_NVX`

## <a href="#_new_built_in_variables" class="anchor"></a>New Built-In Variables

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-positionperview"
  target="_blank" rel="noopener"><code>PositionPerViewNV</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-viewportmaskperview"
  target="_blank" rel="noopener"><code>ViewportMaskPerViewNV</code></a>

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-PerViewAttributesNV"
  target="_blank" rel="noopener"><code>PerViewAttributesNV</code></a>

## <a href="#_examples" class="anchor"></a>Examples

``` c
#version 450 core

#extension GL_KHX_multiview : enable
#extension GL_NVX_multiview_per_view_attributes : enable

layout(location = 0) in vec4 position;
layout(set = 0, binding = 0) uniform Block { mat4 mvpPerView[2]; } buf;

void main()
{
    // Output both per-view positions and gl_Position as a function
    // of gl_ViewIndex
    gl_PositionPerViewNV[0] = buf.mvpPerView[0] * position;
    gl_PositionPerViewNV[1] = buf.mvpPerView[1] * position;
    gl_Position = buf.mvpPerView[gl_ViewIndex] * position;
}
```

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-01-13 (Jeff Bolz)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NVX_multiview_per_view_attributes"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
