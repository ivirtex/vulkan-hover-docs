# VK_GOOGLE_display_timing(3) Manual Page

## Name

VK_GOOGLE_display_timing - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

93

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Ian Elliott <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_GOOGLE_display_timing%5D%20@ianelliottus%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_GOOGLE_display_timing%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>ianelliottus</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-02-14

**IP Status**  
No known IP claims.

**Contributors**  
- Ian Elliott, Google

- Jesse Hall, Google

## <a href="#_description" class="anchor"></a>Description

This device extension allows an application that uses the
[`VK_KHR_swapchain`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html) extension to obtain
information about the presentation engine’s display, to obtain timing
information about each present, and to schedule a present to happen no
earlier than a desired time. An application can use this to minimize
various visual anomalies (e.g. stuttering).

Traditional game and real-time animation applications need to correctly
position their geometry for when the presentable image will be presented
to the user. To accomplish this, applications need various timing
information about the presentation engine’s display. They need to know
when presentable images were actually presented, and when they could
have been presented. Applications also need to tell the presentation
engine to display an image no sooner than a given time. This allows the
application to avoid stuttering, so the animation looks smooth to the
user.

This extension treats variable-refresh-rate (VRR) displays as if they
are fixed-refresh-rate (FRR) displays.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetPastPresentationTimingGOOGLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPastPresentationTimingGOOGLE.html)

- [vkGetRefreshCycleDurationGOOGLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetRefreshCycleDurationGOOGLE.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkPastPresentationTimingGOOGLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPastPresentationTimingGOOGLE.html)

- [VkPresentTimeGOOGLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentTimeGOOGLE.html)

- [VkRefreshCycleDurationGOOGLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRefreshCycleDurationGOOGLE.html)

- Extending [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentInfoKHR.html):

  - [VkPresentTimesInfoGOOGLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentTimesInfoGOOGLE.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_GOOGLE_DISPLAY_TIMING_EXTENSION_NAME`

- `VK_GOOGLE_DISPLAY_TIMING_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PRESENT_TIMES_INFO_GOOGLE`

## <a href="#_examples" class="anchor"></a>Examples

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The example code for the this extension (like the <a
href="VK_KHR_surface.html"><code>VK_KHR_surface</code></a> and
<code>VK_GOOGLE_display_timing</code> extensions) is contained in the
cube demo that is shipped with the official Khronos SDK, and is being
kept up-to-date in that location (see: <a
href="https://github.com/KhronosGroup/Vulkan-Tools/blob/main/cube/cube.c"
class="bare">https://github.com/KhronosGroup/Vulkan-Tools/blob/main/cube/cube.c</a>
).</p></td>
</tr>
</tbody>
</table>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-02-14 (Ian Elliott)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_GOOGLE_display_timing"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
