# VkVideoCapabilitiesKHR(3) Manual Page

## Name

VkVideoCapabilitiesKHR - Structure describing general video capabilities
for a video profile



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoCapabilitiesKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_queue
typedef struct VkVideoCapabilitiesKHR {
    VkStructureType              sType;
    void*                        pNext;
    VkVideoCapabilityFlagsKHR    flags;
    VkDeviceSize                 minBitstreamBufferOffsetAlignment;
    VkDeviceSize                 minBitstreamBufferSizeAlignment;
    VkExtent2D                   pictureAccessGranularity;
    VkExtent2D                   minCodedExtent;
    VkExtent2D                   maxCodedExtent;
    uint32_t                     maxDpbSlots;
    uint32_t                     maxActiveReferencePictures;
    VkExtensionProperties        stdHeaderVersion;
} VkVideoCapabilitiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkVideoCapabilityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilityFlagBitsKHR.html)
  specifying capability flags.

- `minBitstreamBufferOffsetAlignment` is the minimum alignment for
  bitstream buffer offsets.

- `minBitstreamBufferSizeAlignment` is the minimum alignment for
  bitstream buffer range sizes.

- `pictureAccessGranularity` is the granularity at which image access to
  video picture resources happen.

- `minCodedExtent` is the minimum width and height of the coded frames.

- `maxCodedExtent` is the maximum width and height of the coded frames.

- `maxDpbSlots` is the maximum number of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
  target="_blank" rel="noopener">DPB slots</a> supported by a single
  video session.

- `maxActiveReferencePictures` is the maximum number of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#active-reference-pictures"
  target="_blank" rel="noopener">active reference pictures</a> a single
  video coding operation **can** use.

- <span id="video-std-header-version"></span> `stdHeaderVersion` is a
  [VkExtensionProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtensionProperties.html) structure
  reporting the Video Std header name and version supported for the
  video profile.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>It is common for video compression standards to allow using all
reference pictures associated with active DPB slots as active reference
pictures, hence for video decode profiles the values returned in
<code>maxDpbSlots</code> and <code>maxActiveReferencePictures</code> are
often equal. Similarly, in case of video decode profiles supporting
field pictures the value of <code>maxActiveReferencePictures</code>
often equals <code>maxDpbSlots</code> × 2.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-VkVideoCapabilitiesKHR-sType-sType"
  id="VUID-VkVideoCapabilitiesKHR-sType-sType"></a>
  VUID-VkVideoCapabilitiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_CAPABILITIES_KHR`

- <a href="#VUID-VkVideoCapabilitiesKHR-pNext-pNext"
  id="VUID-VkVideoCapabilitiesKHR-pNext-pNext"></a>
  VUID-VkVideoCapabilitiesKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkVideoDecodeAV1CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeAV1CapabilitiesKHR.html),
  [VkVideoDecodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeCapabilitiesKHR.html),
  [VkVideoDecodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264CapabilitiesKHR.html),
  [VkVideoDecodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265CapabilitiesKHR.html),
  [VkVideoEncodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeCapabilitiesKHR.html),
  [VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilitiesKHR.html),
  or
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)

- <a href="#VUID-VkVideoCapabilitiesKHR-sType-unique"
  id="VUID-VkVideoCapabilitiesKHR-sType-unique"></a>
  VUID-VkVideoCapabilitiesKHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkExtensionProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtensionProperties.html),
[VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoCapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilityFlagsKHR.html),
[vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoCapabilitiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
