# VK_KHR_display_swapchain(3) Manual Page

## Name

VK_KHR_display_swapchain - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

4

## <a href="#_revision" class="anchor"></a>Revision

10

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html)  
and  
[VK_KHR_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_display.html)  

## <a href="#_contact" class="anchor"></a>Contact

- James Jones <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_display_swapchain%5D%20@cubanismo%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_display_swapchain%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>cubanismo</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-03-13

**IP Status**  
No known IP claims.

**Contributors**  
- James Jones, NVIDIA

- Jeff Vigil, Qualcomm

- Jesse Hall, Google

## <a href="#_description" class="anchor"></a>Description

This extension provides an API to create a swapchain directly on a
device’s display without any underlying window system.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCreateSharedSwapchainsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSharedSwapchainsKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentInfoKHR.html):

  - [VkDisplayPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPresentInfoKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_DISPLAY_SWAPCHAIN_EXTENSION_NAME`

- `VK_KHR_DISPLAY_SWAPCHAIN_SPEC_VERSION`

- Extending [VkResult](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html):

  - `VK_ERROR_INCOMPATIBLE_DISPLAY_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_DISPLAY_PRESENT_INFO_KHR`

## <a href="#_issues" class="anchor"></a>Issues

1\) Should swapchains sharing images each hold a reference to the
images, or should it be up to the application to destroy the swapchains
and images in an order that avoids the need for reference counting?

**RESOLVED**: Take a reference. The lifetime of presentable images is
already complex enough.

2\) Should the `srcRect` and `dstRect` parameters be specified as part
of the presentation command, or at swapchain creation time?

**RESOLVED**: As part of the presentation command. This allows moving
and scaling the image on the screen without the need to respecify the
mode or create a new swapchain and presentable images.

3\) Should `srcRect` and `dstRect` be specified as rects, or separate
offset/extent values?

**RESOLVED**: As rects. Specifying them separately might make it easier
for hardware to expose support for one but not the other, but in such
cases applications must just take care to obey the reported capabilities
and not use non-zero offsets or extents that require scaling, as
appropriate.

4\) How can applications create multiple swapchains that use the same
images?

**RESOLVED**: By calling
[vkCreateSharedSwapchainsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSharedSwapchainsKHR.html).

An earlier resolution used
[vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSwapchainKHR.html), chaining multiple
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html) structures
through `pNext`. In order to allow each swapchain to also allow other
extension structs, a level of indirection was used:
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html)::`pNext`
pointed to a different structure, which had both `sType` and `pNext`
members for additional extensions, and also had a pointer to the next
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html) structure. The
number of swapchains to be created could only be found by walking this
linked list of alternating structures, and the `pSwapchains` out
parameter was reinterpreted to be an array of
[VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html) handles.

Another option considered was a method to specify a “shared” swapchain
when creating a new swapchain, such that groups of swapchains using the
same images could be built up one at a time. This was deemed unusable
because drivers need to know all of the displays an image will be used
on when determining which internal formats and layouts to use for that
image.

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
<p>The example code for the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_display.html"><code>VK_KHR_display</code></a> and
<code>VK_KHR_display_swapchain</code> extensions was removed from the
appendix after revision 1.0.43. The display swapchain creation example
code was ported to the cube demo that is shipped with the official
Khronos SDK, and is being kept up-to-date in that location (see: <a
href="https://github.com/KhronosGroup/Vulkan-Tools/blob/main/cube/cube.c"
class="bare">https://github.com/KhronosGroup/Vulkan-Tools/blob/main/cube/cube.c</a>).</p></td>
</tr>
</tbody>
</table>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2015-07-29 (James Jones)

  - Initial draft

- Revision 2, 2015-08-21 (Ian Elliott)

  - Renamed this extension and all of its enumerations, types,
    functions, etc. This makes it compliant with the proposed standard
    for Vulkan extensions.

  - Switched from “revision” to “version”, including use of the
    VK_MAKE_VERSION macro in the header file.

- Revision 3, 2015-09-01 (James Jones)

  - Restore single-field revision number.

- Revision 4, 2015-09-08 (James Jones)

  - Allow creating multiple swapchains that share the same images using
    a single call to vkCreateSwapchainKHR().

- Revision 5, 2015-09-10 (Alon Or-bach)

  - Removed underscores from SWAP_CHAIN in two enums.

- Revision 6, 2015-10-02 (James Jones)

  - Added support for smart panels/buffered displays.

- Revision 7, 2015-10-26 (Ian Elliott)

  - Renamed from VK_EXT_KHR_display_swapchain to
    VK_KHR_display_swapchain.

- Revision 8, 2015-11-03 (Daniel Rakos)

  - Updated sample code based on the changes to VK_KHR_swapchain.

- Revision 9, 2015-11-10 (Jesse Hall)

  - Replaced VkDisplaySwapchainCreateInfoKHR with
    vkCreateSharedSwapchainsKHR, changing resolution of issue \#4.

- Revision 10, 2017-03-13 (James Jones)

  - Closed all remaining issues. The specification and implementations
    have been shipping with the proposed resolutions for some time now.

  - Removed the sample code and noted it has been integrated into the
    official Vulkan SDK cube demo.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_display_swapchain"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
