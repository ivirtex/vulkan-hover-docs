# VK\_NVX\_multiview\_per\_view\_attributes(3) Manual Page

## Name

VK\_NVX\_multiview\_per\_view\_attributes - device extension



## [](#_registered_extension_number)Registered Extension Number

98

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_multiview](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_multiview.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_api_interactions)API Interactions

- Interacts with VK\_VERSION\_1\_3
- Interacts with VK\_KHR\_dynamic\_rendering

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_NVX\_multiview\_per\_view\_attributes](https://github.khronos.org/SPIRV-Registry/extensions/NVX/SPV_NVX_multiview_per_view_attributes.html)

## [](#_contact)Contact

- Jeff Bolz [\[GitHub\]jeffbolznv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NVX_multiview_per_view_attributes%5D%20%40jeffbolznv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NVX_multiview_per_view_attributes%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-01-13

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- This extension provides API support for [`GL_NVX_multiview_per_view_attributes`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nvx/GL_NVX_multiview_per_view_attributes.txt)
- This extension interacts with `VK_NV_viewport_array2`.

**Contributors**

- Jeff Bolz, NVIDIA
- Daniel Koch, NVIDIA

## [](#_description)Description

This extension adds a new way to write shaders to be used with multiview subpasses, where the attributes for all views are written out by a single invocation of the [pre-rasterization shader stages](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization). Related SPIR-V and GLSL extensions `SPV_NVX_multiview_per_view_attributes` and `GL_NVX_multiview_per_view_attributes` introduce per-view position and viewport mask attributes arrays, and this extension defines how those per-view attribute arrays are interpreted by Vulkan. Pipelines using per-view attributes **may** only execute the [pre-rasterization shader stages](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization) once for all views rather than once per-view, which reduces redundant shading work.

A subpass creation flag controls whether the subpass uses this extension. A subpass **must** either exclusively use this extension or not use it at all.

Some Vulkan implementations only support the position attribute varying between views in the X component. A subpass can declare via a second creation flag whether all pipelines compiled for this subpass will obey this restriction.

Shaders that use the new per-view outputs (e.g. `gl_PositionPerViewNV`) **must** also write the non-per-view output (`gl_Position`), and the values written **must** be such that `gl_Position = gl_PositionPerViewNV[gl_ViewIndex]` for all views in the subpass. Implementations are free to either use the per-view outputs or the non-per-view outputs, whichever would be more efficient.

If the `VK_NV_viewport_array2` extension is not also supported and enabled, the per-view viewport mask **must** not be used.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX.html)

If [Vulkan Version 1.3](#versions-1.3) or [VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html) is supported:

- Extending [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceInfo.html), [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html):
  
  - [VkMultiviewPerViewAttributesInfoNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMultiviewPerViewAttributesInfoNVX.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NVX_MULTIVIEW_PER_VIEW_ATTRIBUTES_EXTENSION_NAME`
- `VK_NVX_MULTIVIEW_PER_VIEW_ATTRIBUTES_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PER_VIEW_ATTRIBUTES_PROPERTIES_NVX`
- Extending [VkSubpassDescriptionFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescriptionFlagBits.html):
  
  - `VK_SUBPASS_DESCRIPTION_PER_VIEW_ATTRIBUTES_BIT_NVX`
  - `VK_SUBPASS_DESCRIPTION_PER_VIEW_POSITION_X_ONLY_BIT_NVX`

If [Vulkan Version 1.3](#versions-1.3) or [VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html) is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_MULTIVIEW_PER_VIEW_ATTRIBUTES_INFO_NVX`

## [](#_new_built_in_variables)New Built-In Variables

- [`PositionPerViewNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-positionperview)
- [`ViewportMaskPerViewNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-viewportmaskperview)

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`PerViewAttributesNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-PerViewAttributesNV)

## [](#_examples)Examples

```c
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

## [](#_version_history)Version History

- Revision 1, 2017-01-13 (Jeff Bolz)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NVX_multiview_per_view_attributes).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0