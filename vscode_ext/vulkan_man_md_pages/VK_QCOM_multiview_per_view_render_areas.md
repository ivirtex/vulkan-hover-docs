# VK_QCOM_multiview_per_view_render_areas(3) Manual Page

## Name

VK_QCOM_multiview_per_view_render_areas - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

511

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_contact" class="anchor"></a>Contact

- Matthew Netsch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QCOM_multiview_per_view_render_areas%5D%20@mnetsch%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QCOM_multiview_per_view_render_areas%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>mnetsch</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-01-10

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- This extension interacts with
  [`VK_KHR_dynamic_rendering`](VK_KHR_dynamic_rendering.html)

- This extension interacts with
  [`VK_QCOM_render_pass_transform`](VK_QCOM_render_pass_transform.html)

**Contributors**  
- Jeff Leger, Qualcomm

- Jonathan Tinkham, Qualcomm

- Jonathan Wicks, Qualcomm

## <a href="#_description" class="anchor"></a>Description

Certain use cases (e.g., side-by-side VR rendering) use multiview and
render to distinct regions of the framebuffer for each view. On some
implementations, there may be a performance benefit for providing
per-view render areas to the implementation. Such per-view render areas
can be used by the implementation to reduce the pixels that are affected
by attachment load, store, and multisample resolve operations.

The extension enables a multiview render pass instance to define
per-view render areas. For each view of a multiview render pass
instance, only those pixels in the per-view render area are affected by
load, store and resolve operations.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceMultiviewPerViewRenderAreasFeaturesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMultiviewPerViewRenderAreasFeaturesQCOM.html)

- Extending [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html),
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html):

  - [VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_QCOM_MULTIVIEW_PER_VIEW_RENDER_AREAS_EXTENSION_NAME`

- `VK_QCOM_MULTIVIEW_PER_VIEW_RENDER_AREAS_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_MULTIVIEW_PER_VIEW_RENDER_AREAS_RENDER_PASS_BEGIN_INFO_QCOM`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PER_VIEW_RENDER_AREAS_FEATURES_QCOM`

## <a href="#_issues" class="anchor"></a>Issues

1\) Do the per-view `renderAreas` interact with
[vkGetRenderAreaGranularity](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetRenderAreaGranularity.html) ?

**RESOLVED**: There is no change. The granularity returned by
[vkGetRenderAreaGranularity](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetRenderAreaGranularity.html) also
applies to the per-view `renderAreas`.

2\) How does this extension interact with
[`VK_QCOM_render_pass_transform`](VK_QCOM_render_pass_transform.html)?

**RESOLVED**: When
[`VK_QCOM_render_pass_transform`](VK_QCOM_render_pass_transform.html) is
enabled, the application provides render area in non-rotated coordinates
which is rotated by the implementation to the rotated coordinate system.
When this extension is used in combination with
[`VK_QCOM_render_pass_transform`](VK_QCOM_render_pass_transform.html),
then the `renderArea` provided in
[VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`renderArea`,
[VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html)::`renderArea`, or
[VkCommandBufferInheritanceRenderPassTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceRenderPassTransformInfoQCOM.html)::`renderArea`
is rotated by the implementation. The per-view render areas are not
rotated.

3\) How does this extension interact with
[`VK_QCOM_multiview_per_view_viewports`](VK_QCOM_multiview_per_view_viewports.html)

**RESOLVED**: There is no direct interaction. The per-view viewports and
the per-view renderAreas are orthogonal features.

4\) When a per-view `renderArea` is specified, must multiview rendering
for each view of a multiview render pass be contained within the
per-view `renderArea`?

**RESOLVED**: Yes, and the
[`VK_QCOM_multiview_per_view_viewports`](VK_QCOM_multiview_per_view_viewports.html)
may help here since it provides per-view scissors.

5\) When per-view render areas are specified, what purpose if any do
[VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html)::`renderArea` and
[VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`renderArea` serve?

**RESOLVED**: The per-view `renderArea` effectively overrides the
per-renderpass `renderArea`. The per-view `renderArea` defines the
regions of the attachments that are effected by load, store, and
multisample resolve operations. A valid implementation could ignore the
per-renderpass `renderArea`. However, as an aid to the implementation,
the application must set the per-renderpass `renderArea` to an area that
is at least as large as the union of all the per-view render areas.
Pixels that are within the per-renderpass `renderArea` but not within
any per-view render area must not be affected by load, store, or
multisample resolve operations.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-01-10 (Jeff Leger)

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_QCOM_multiview_per_view_render_areas"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
