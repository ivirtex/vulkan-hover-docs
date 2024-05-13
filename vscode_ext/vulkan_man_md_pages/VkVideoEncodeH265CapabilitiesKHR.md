# VkVideoEncodeH265CapabilitiesKHR(3) Manual Page

## Name

VkVideoEncodeH265CapabilitiesKHR - Structure describing H.265 encode
capabilities



## <a href="#_c_specification" class="anchor"></a>C Specification

When calling
[vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
to query the capabilities for an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-profile"
target="_blank" rel="noopener">H.265 encode profile</a>, the
[VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html)::`pNext` chain
**must** include a `VkVideoEncodeH265CapabilitiesKHR` structure that
will be filled with the profile-specific capabilities.

The `VkVideoEncodeH265CapabilitiesKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_encode_h265
typedef struct VkVideoEncodeH265CapabilitiesKHR {
    VkStructureType                                sType;
    void*                                          pNext;
    VkVideoEncodeH265CapabilityFlagsKHR            flags;
    StdVideoH265LevelIdc                           maxLevelIdc;
    uint32_t                                       maxSliceSegmentCount;
    VkExtent2D                                     maxTiles;
    VkVideoEncodeH265CtbSizeFlagsKHR               ctbSizes;
    VkVideoEncodeH265TransformBlockSizeFlagsKHR    transformBlockSizes;
    uint32_t                                       maxPPictureL0ReferenceCount;
    uint32_t                                       maxBPictureL0ReferenceCount;
    uint32_t                                       maxL1ReferenceCount;
    uint32_t                                       maxSubLayerCount;
    VkBool32                                       expectDyadicTemporalSubLayerPattern;
    int32_t                                        minQp;
    int32_t                                        maxQp;
    VkBool32                                       prefersGopRemainingFrames;
    VkBool32                                       requiresGopRemainingFrames;
    VkVideoEncodeH265StdFlagsKHR                   stdSyntaxFlags;
} VkVideoEncodeH265CapabilitiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkVideoEncodeH265CapabilityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilityFlagBitsKHR.html)
  indicating supported H.265 encoding capabilities.

- `maxLevelIdc` is a `StdVideoH265LevelIdc` value indicating the maximum
  H.265 level supported by the profile, where enum constant
  `STD_VIDEO_H265_LEVEL_IDC_<major>_<minor>` identifies H.265 level
  `<major>.<minor>` as defined in section A.4 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>.

- `maxSliceSegmentCount` indicates the maximum number of slice segments
  that **can** be encoded for a single picture. Further restrictions
  **may** apply to the number of slice segments that **can** be encoded
  for a single picture depending on other capabilities and
  codec-specific rules.

- `maxTiles` indicates the maximum number of H.265 tile columns and
  rows, as defined in sections 3.175 and 3.176 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a> that
  **can** be encoded for a single picture. Further restrictions **may**
  apply to the number of H.265 tiles that **can** be encoded for a
  single picture depending on other capabilities and codec-specific
  rules.

- `ctbSizes` is a bitmask of
  [VkVideoEncodeH265CtbSizeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CtbSizeFlagBitsKHR.html)
  describing the supported CTB sizes.

- `transformBlockSizes` is a bitmask of
  [VkVideoEncodeH265TransformBlockSizeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265TransformBlockSizeFlagBitsKHR.html)
  describing the supported transform block sizes.

- `maxPPictureL0ReferenceCount` indicates the maximum number of
  reference pictures the implementation supports in the reference list
  L0 for <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-p-pic"
  target="_blank" rel="noopener">P pictures</a>.

  <table>
  <colgroup>
  <col style="width: 50%" />
  <col style="width: 50%" />
  </colgroup>
  <tbody>
  <tr class="odd">
  <td class="icon"><em></em></td>
  <td class="content">Note
  <p>As implementations <strong>may</strong> <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-overrides"
  target="_blank" rel="noopener">override</a> the reference lists,
  <code>maxPPictureL0ReferenceCount</code> does not limit the number of
  elements that the application <strong>can</strong> specify in the L0
  reference list for P pictures. However, if
  <code>maxPPictureL0ReferenceCount</code> is zero, then the use of P
  pictures is not allowed. In case of H.265 encoding, backward-only
  predictive pictures <strong>can</strong> be encoded even if P pictures
  are not supported, as the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a> supports
  <em>generalized P &amp; B frames</em> (also known as low delay B frames)
  whereas B frames <strong>can</strong> refer to past frames through both
  the L0 and L1 reference lists.</p></td>
  </tr>
  </tbody>
  </table>

- `maxBPictureL0ReferenceCount` indicates the maximum number of
  reference pictures the implementation supports in the reference list
  L0 for <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-b-pic"
  target="_blank" rel="noopener">B pictures</a>.

- `maxL1ReferenceCount` indicates the maximum number of reference
  pictures the implementation supports in the reference list L1 if
  encoding of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-b-pic"
  target="_blank" rel="noopener">B pictures</a> is supported.

  <table>
  <colgroup>
  <col style="width: 50%" />
  <col style="width: 50%" />
  </colgroup>
  <tbody>
  <tr class="odd">
  <td class="icon"><em></em></td>
  <td class="content">Note
  <p>As implementations <strong>may</strong> <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-overrides"
  target="_blank" rel="noopener">override</a> the reference lists,
  <code>maxBPictureL0ReferenceCount</code> and
  <code>maxL1ReferenceCount</code> does not limit the number of elements
  that the application <strong>can</strong> specify in the L0 and L1
  reference lists for B pictures. However, if
  <code>maxBPictureL0ReferenceCount</code> and
  <code>maxL1ReferenceCount</code> are both zero, then the use of B
  pictures is not allowed.</p></td>
  </tr>
  </tbody>
  </table>

- `maxSubLayerCount` indicates the maximum number of H.265 sub-layers
  supported by the implementation.

- `expectDyadicTemporalSubLayerPattern` indicates that the
  implementation’s rate control algorithms expect the application to use
  a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-layer-pattern-dyadic"
  target="_blank" rel="noopener">dyadic temporal sub-layer pattern</a>
  when encoding multiple temporal sub-layers.

- `minQp` indicates the minimum QP value supported.

- `maxQp` indicates the maximum QP value supported.

- `prefersGopRemainingFrames` indicates that the implementation’s rate
  control algorithm prefers the application to specify the number of
  frames of each type <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-gop-remaining-frames"
  target="_blank" rel="noopener">remaining</a> in the current <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-gop"
  target="_blank" rel="noopener">group of pictures</a> when beginning a
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-coding-scope"
  target="_blank" rel="noopener">video coding scope</a>.

- `requiresGopRemainingFrames` indicates that the implementation’s rate
  control algorithm requires the application to specify the number of
  frames of each type <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-gop-remaining-frames"
  target="_blank" rel="noopener">remaining</a> in the current <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-gop"
  target="_blank" rel="noopener">group of pictures</a> when beginning a
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-coding-scope"
  target="_blank" rel="noopener">video coding scope</a>.

- `stdSyntaxFlags` is a bitmask of
  [VkVideoEncodeH265StdFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265StdFlagBitsKHR.html)
  indicating capabilities related to H.265 syntax elements.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkVideoEncodeH265CapabilitiesKHR-sType-sType"
  id="VUID-VkVideoEncodeH265CapabilitiesKHR-sType-sType"></a>
  VUID-VkVideoEncodeH265CapabilitiesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_CAPABILITIES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_h265](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_h265.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoEncodeH265CapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilityFlagsKHR.html),
[VkVideoEncodeH265CtbSizeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CtbSizeFlagsKHR.html),
[VkVideoEncodeH265StdFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265StdFlagsKHR.html),
[VkVideoEncodeH265TransformBlockSizeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265TransformBlockSizeFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeH265CapabilitiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
