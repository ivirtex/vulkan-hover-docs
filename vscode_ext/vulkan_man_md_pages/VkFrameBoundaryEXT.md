# VkFrameBoundaryEXT(3) Manual Page

## Name

VkFrameBoundaryEXT - Add frame boundary information to queue submissions



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkFrameBoundaryEXT` structure is defined as:

``` c
// Provided by VK_EXT_frame_boundary
typedef struct VkFrameBoundaryEXT {
    VkStructureType            sType;
    const void*                pNext;
    VkFrameBoundaryFlagsEXT    flags;
    uint64_t                   frameID;
    uint32_t                   imageCount;
    const VkImage*             pImages;
    uint32_t                   bufferCount;
    const VkBuffer*            pBuffers;
    uint64_t                   tagName;
    size_t                     tagSize;
    const void*                pTag;
} VkFrameBoundaryEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkFrameBoundaryFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFrameBoundaryFlagBitsEXT.html) that can
  flag the last submission of a frame identifier.

- `frameID` is the frame identifier.

- `imageCount` is the number of images that store frame results.

- `pImages` is a pointer to an array of VkImage objects with imageCount
  entries.

- `bufferCount` is the number of buffers the store the frame results.

- `pBuffers` is a pointer to an array of VkBuffer objects with
  bufferCount entries.

- `tagName` is a numerical identifier for tag data.

- `tagSize` is the number of bytes of tag data.

- `pTag` is a pointer to an array of `tagSize` bytes containing tag
  data.

## <a href="#_description" class="anchor"></a>Description

The application **can** associate frame boundary information to a queue
submission call by adding a `VkFrameBoundaryEXT` structure to the
`pNext` chain of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-submission"
target="_blank" rel="noopener">queue submission</a>,
[VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentInfoKHR.html), or
[VkBindSparseInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindSparseInfo.html).

The frame identifier is used to associate one or more queue submission
to a frame, it is thus meant to be unique within a frame lifetime, i.e.
it is possible (but not recommended) to reuse frame identifiers, as long
as any two frames with any chance of having overlapping queue
submissions (as in the example above) use two different frame
identifiers.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Since the concept of frame is application-dependent, there is no way
to validate the use of frame identifier. It is good practice to use a
monotonically increasing counter as the frame identifier and not reuse
identifiers between frames.</p></td>
</tr>
</tbody>
</table>

The `pImages` and `pBuffers` arrays contain a list of images and buffers
which store the "end result" of the frame. As the concept of frame is
application-dependent, not all frames **may** produce their results in
images or buffers, yet this is a sufficiently common case to be handled
by `VkFrameBoundaryEXT`. Note that no extra information, such as image
layout is being provided, since the images are meant to be used by tools
which would already be tracking this required information. Having the
possibility of passing a list of end-result images makes
`VkFrameBoundaryEXT` as expressive as
[vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueuePresentKHR.html), which is often the default
frame boundary delimiter.

The application **can** also associate arbitrary extra information via
tag data using `tagName`, `tagSize` and `pTag`. This extra information
is typically tool-specific.

Valid Usage (Implicit)

- <a href="#VUID-VkFrameBoundaryEXT-sType-sType"
  id="VUID-VkFrameBoundaryEXT-sType-sType"></a>
  VUID-VkFrameBoundaryEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_FRAME_BOUNDARY_EXT`

- <a href="#VUID-VkFrameBoundaryEXT-flags-parameter"
  id="VUID-VkFrameBoundaryEXT-flags-parameter"></a>
  VUID-VkFrameBoundaryEXT-flags-parameter  
  `flags` **must** be a valid combination of
  [VkFrameBoundaryFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFrameBoundaryFlagBitsEXT.html) values

- <a href="#VUID-VkFrameBoundaryEXT-pImages-parameter"
  id="VUID-VkFrameBoundaryEXT-pImages-parameter"></a>
  VUID-VkFrameBoundaryEXT-pImages-parameter  
  If `imageCount` is not `0`, and `pImages` is not `NULL`, `pImages`
  **must** be a valid pointer to an array of `imageCount` valid
  [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handles

- <a href="#VUID-VkFrameBoundaryEXT-pBuffers-parameter"
  id="VUID-VkFrameBoundaryEXT-pBuffers-parameter"></a>
  VUID-VkFrameBoundaryEXT-pBuffers-parameter  
  If `bufferCount` is not `0`, and `pBuffers` is not `NULL`, `pBuffers`
  **must** be a valid pointer to an array of `bufferCount` valid
  [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handles

- <a href="#VUID-VkFrameBoundaryEXT-pTag-parameter"
  id="VUID-VkFrameBoundaryEXT-pTag-parameter"></a>
  VUID-VkFrameBoundaryEXT-pTag-parameter  
  If `tagSize` is not `0`, and `pTag` is not `NULL`, `pTag` **must** be
  a valid pointer to an array of `tagSize` bytes

- <a href="#VUID-VkFrameBoundaryEXT-commonparent"
  id="VUID-VkFrameBoundaryEXT-commonparent"></a>
  VUID-VkFrameBoundaryEXT-commonparent  
  Both of the elements of `pBuffers`, and the elements of `pImages` that
  are valid handles of non-ignored parameters **must** have been
  created, allocated, or retrieved from the same
  [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_frame_boundary](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_frame_boundary.html),
[VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html),
[VkFrameBoundaryFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFrameBoundaryFlagsEXT.html),
[VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkFrameBoundaryEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
