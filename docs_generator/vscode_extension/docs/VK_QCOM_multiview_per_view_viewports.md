# VK\_QCOM\_multiview\_per\_view\_viewports(3) Manual Page

## Name

VK\_QCOM\_multiview\_per\_view\_viewports - device extension



## [](#_registered_extension_number)Registered Extension Number

489

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Matthew Netsch [\[GitHub\]mnetsch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QCOM_multiview_per_view_viewports%5D%20%40mnetsch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QCOM_multiview_per_view_viewports%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2022-11-22

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- This extension interacts with `VK_KHR_dynamic_rendering`
- This extension interacts with `VK_EXT_extended_dynamic_state`

**Contributors**

- Jeff Leger, Qualcomm
- Jonathan Tinkham, Qualcomm
- Jonathan Wicks, Qualcomm

## [](#_description)Description

Certain use cases for multiview have a need for specifying a separate viewport and scissor for each view, without using shader-based viewport indexing as introduced with `VK_EXT_shader_viewport_index_layer`.

This extension adds a new way to control ViewportIndex with multiview. When the [`multiviewPerViewViewports`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-multiviewPerViewViewports) feature is enabled and if the last pre-rasterization shader entry point’s interface does not use the `ViewportIndex` built-in decoration, then each view of a multiview render pass instance will use a viewport and scissor index equal to the `ViewIndex`.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceMultiviewPerViewViewportsFeaturesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMultiviewPerViewViewportsFeaturesQCOM.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_QCOM_MULTIVIEW_PER_VIEW_VIEWPORTS_EXTENSION_NAME`
- `VK_QCOM_MULTIVIEW_PER_VIEW_VIEWPORTS_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PER_VIEW_VIEWPORTS_FEATURES_QCOM`

## [](#_issues)Issues

1\) Is it possible to enable/disable the [`multiviewPerViewViewports`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-multiviewPerViewViewports) feature for individual render pass instances?

**RESOLVED**: No, when the multiviewPerViewViewports feature is enabled during vkCreateDevice, then all created render pass instances (including dynamic render passes from `VK_KHR_dynamic_rendering`) and all created VkPipelines will have the feature enabled. This approach was chosen because it simplifies application code and there is no known use case to enable/disable the feature for individual render passes or pipelines.

2\) When this extension is used, is the value of `ViewportIndex` implicitly written by the last pre-rasterization shader stage and can the value of `ViewportIndex` be read in the fragment shader?

**RESOLVED**: No, use of the extension does not add an implicit write to `ViewportIndex` in any shader stage, and additionally, the value of `ViewportIndex` in the fragment shader is undefined.

## [](#_version_history)Version History

- Revision 1, 2022-11-22 (Jeff Leger)

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_QCOM_multiview_per_view_viewports)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0