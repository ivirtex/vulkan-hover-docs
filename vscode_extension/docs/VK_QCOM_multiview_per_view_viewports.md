# VK_QCOM_multiview_per_view_viewports(3) Manual Page

## Name

VK_QCOM_multiview_per_view_viewports - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

489

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Matthew Netsch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QCOM_multiview_per_view_viewports%5D%20@mnetsch%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QCOM_multiview_per_view_viewports%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>mnetsch</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2022-11-22

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- This extension interacts with
  [`VK_KHR_dynamic_rendering`](VK_KHR_dynamic_rendering.html)

- This extension interacts with
  [`VK_EXT_extended_dynamic_state`](VK_EXT_extended_dynamic_state.html)

**Contributors**  
- Jeff Leger, Qualcomm

- Jonathan Tinkham, Qualcomm

- Jonathan Wicks, Qualcomm

## <a href="#_description" class="anchor"></a>Description

Certain use cases for multiview have a need for specifying a separate
viewport and scissor for each view, without using shader-based viewport
indexing as introduced with
[`VK_EXT_shader_viewport_index_layer`](VK_EXT_shader_viewport_index_layer.html).

This extension adds a new way to control ViewportIndex with multiview.
When the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiview-per-view-viewports"
target="_blank"
rel="noopener"><code>multiviewPerViewViewports</code></a> feature is
enabled and if the last pre-rasterization shader entry point’s interface
does not use the `ViewportIndex` built-in decoration, then each view of
a multiview render pass instance will use a viewport and scissor index
equal to the `ViewIndex`.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceMultiviewPerViewViewportsFeaturesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMultiviewPerViewViewportsFeaturesQCOM.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_QCOM_MULTIVIEW_PER_VIEW_VIEWPORTS_EXTENSION_NAME`

- `VK_QCOM_MULTIVIEW_PER_VIEW_VIEWPORTS_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PER_VIEW_VIEWPORTS_FEATURES_QCOM`

## <a href="#_issues" class="anchor"></a>Issues

1\) Is it possible to enable/disable the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiview-per-view-viewports"
target="_blank"
rel="noopener"><code>multiviewPerViewViewports</code></a> feature for
individual render pass instances?

**RESOLVED**: No, when the multiviewPerViewViewports feature is enabled
during vkCreateDevice, then all created render pass instances (including
dynamic render passes from
[`VK_KHR_dynamic_rendering`](VK_KHR_dynamic_rendering.html)) and all
created VkPipelines will have the feature enabled. This approach was
chosen because it simplifies application code and there is no known use
case to enable/disable the feature for individual render passes or
pipelines.

2\) When this extension is used, is the value of `ViewportIndex`
implicitly written by the last pre-rasterization shader stage and can
the value of `ViewportIndex` be read in the fragment shader?

**RESOLVED**: No, use of the extension does not add an implicit write to
`ViewportIndex` in any shader stage, and additionally, the value of
`ViewportIndex` in the fragment shader is undefined.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2022-11-22 (Jeff Leger)

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_QCOM_multiview_per_view_viewports"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
