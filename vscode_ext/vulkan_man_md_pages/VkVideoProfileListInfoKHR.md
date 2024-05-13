# VkVideoProfileListInfoKHR(3) Manual Page

## Name

VkVideoProfileListInfoKHR - Structure specifying one or more video
profiles used in conjunction



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoProfileListInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_queue
typedef struct VkVideoProfileListInfoKHR {
    VkStructureType                 sType;
    const void*                     pNext;
    uint32_t                        profileCount;
    const VkVideoProfileInfoKHR*    pProfiles;
} VkVideoProfileListInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `profileCount` is the number of elements in the `pProfiles` array.

- `pProfiles` is a pointer to an array of
  [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html) structures.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Video transcoding is an example of a use case that necessitates the
specification of multiple profiles in various contexts.</p></td>
</tr>
</tbody>
</table>

When the application provides a video decode profile and one or more
video encode profiles in the profile list, the implementation ensures
that any capabilitities returned or resources created are suitable for
the video transcoding use cases without the need for manual data
transformations.

Valid Usage

- <a href="#VUID-VkVideoProfileListInfoKHR-pProfiles-06813"
  id="VUID-VkVideoProfileListInfoKHR-pProfiles-06813"></a>
  VUID-VkVideoProfileListInfoKHR-pProfiles-06813  
  `pProfiles` **must** not contain more than one element whose
  `videoCodecOperation` member specifies a decode operation

Valid Usage (Implicit)

- <a href="#VUID-VkVideoProfileListInfoKHR-sType-sType"
  id="VUID-VkVideoProfileListInfoKHR-sType-sType"></a>
  VUID-VkVideoProfileListInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_PROFILE_LIST_INFO_KHR`

- <a href="#VUID-VkVideoProfileListInfoKHR-pProfiles-parameter"
  id="VUID-VkVideoProfileListInfoKHR-pProfiles-parameter"></a>
  VUID-VkVideoProfileListInfoKHR-pProfiles-parameter  
  If `profileCount` is not `0`, `pProfiles` **must** be a valid pointer
  to an array of `profileCount` valid
  [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html) structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoProfileListInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
