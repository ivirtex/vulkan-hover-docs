# VkVideoEncodeH264CapabilitiesKHR(3) Manual Page

## Name

VkVideoEncodeH264CapabilitiesKHR - Structure describing H.264 encode
capabilities



## <a href="#_c_specification" class="anchor"></a>C Specification

When calling
[vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
to query the capabilities for an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-profile"
target="_blank" rel="noopener">H.264 encode profile</a>, the
[VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html)::`pNext` chain
**must** include a `VkVideoEncodeH264CapabilitiesKHR` structure that
will be filled with the profile-specific capabilities.

The `VkVideoEncodeH264CapabilitiesKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_encode_h264
typedef struct VkVideoEncodeH264CapabilitiesKHR {
    VkStructureType                        sType;
    void*                                  pNext;
    VkVideoEncodeH264CapabilityFlagsKHR    flags;
    StdVideoH264LevelIdc                   maxLevelIdc;
    uint32_t                               maxSliceCount;
    uint32_t                               maxPPictureL0ReferenceCount;
    uint32_t                               maxBPictureL0ReferenceCount;
    uint32_t                               maxL1ReferenceCount;
    uint32_t                               maxTemporalLayerCount;
    VkBool32                               expectDyadicTemporalLayerPattern;
    int32_t                                minQp;
    int32_t                                maxQp;
    VkBool32                               prefersGopRemainingFrames;
    VkBool32                               requiresGopRemainingFrames;
    VkVideoEncodeH264StdFlagsKHR           stdSyntaxFlags;
} VkVideoEncodeH264CapabilitiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkVideoEncodeH264CapabilityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilityFlagBitsKHR.html)
  indicating supported H.264 encoding capabilities.

- `maxLevelIdc` is a `StdVideoH264LevelIdc` value indicating the maximum
  H.264 level supported by the profile, where enum constant
  `STD_VIDEO_H264_LEVEL_IDC_<major>_<minor>` identifies H.264 level
  `<major>.<minor>` as defined in section A.3 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
  target="_blank" rel="noopener">ITU-T H.264 Specification</a>.

- `maxSliceCount` indicates the maximum number of slices that **can** be
  encoded for a single picture. Further restrictions **may** apply to
  the number of slices that **can** be encoded for a single picture
  depending on other capabilities and codec-specific rules.

- `maxPPictureL0ReferenceCount` indicates the maximum number of
  reference pictures the implementation supports in the reference list
  L0 for <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-p-pic"
  target="_blank" rel="noopener">P pictures</a>.

  <table>
  <colgroup>
  <col style="width: 50%" />
  <col style="width: 50%" />
  </colgroup>
  <tbody>
  <tr>
  <td class="icon"><em></em></td>
  <td class="content">Note
  <p>As implementations <strong>may</strong> <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-overrides"
  target="_blank" rel="noopener">override</a> the reference lists,
  <code>maxPPictureL0ReferenceCount</code> does not limit the number of
  elements that the application <strong>can</strong> specify in the L0
  reference list for P pictures. However, if
  <code>maxPPictureL0ReferenceCount</code> is zero, then the use of P
  pictures is not allowed.</p></td>
  </tr>
  </tbody>
  </table>

- `maxBPictureL0ReferenceCount` indicates the maximum number of
  reference pictures the implementation supports in the reference list
  L0 for <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-b-pic"
  target="_blank" rel="noopener">B pictures</a>.

- `maxL1ReferenceCount` indicates the maximum number of reference
  pictures the implementation supports in the reference list L1 if
  encoding of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-b-pic"
  target="_blank" rel="noopener">B pictures</a> is supported.

  <table>
  <colgroup>
  <col style="width: 50%" />
  <col style="width: 50%" />
  </colgroup>
  <tbody>
  <tr>
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

- `maxTemporalLayerCount` indicates the maximum number of H.264 temporal
  layers supported by the implementation.

- `expectDyadicTemporalLayerPattern` indicates that the implementation’s
  rate control algorithms expect the application to use a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-layer-pattern-dyadic"
  target="_blank" rel="noopener">dyadic temporal layer pattern</a> when
  encoding multiple temporal layers.

- `minQp` indicates the minimum QP value supported.

- `maxQp` indicates the maximum QP value supported.

- `prefersGopRemainingFrames` indicates that the implementation’s rate
  control algorithm prefers the application to specify the number of
  frames of each type <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-gop-remaining-frames"
  target="_blank" rel="noopener">remaining</a> in the current <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-gop"
  target="_blank" rel="noopener">group of pictures</a> when beginning a
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-coding-scope"
  target="_blank" rel="noopener">video coding scope</a>.

- `requiresGopRemainingFrames` indicates that the implementation’s rate
  control algorithm requires the application to specify the number of
  frames of each type <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-gop-remaining-frames"
  target="_blank" rel="noopener">remaining</a> in the current <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-gop"
  target="_blank" rel="noopener">group of pictures</a> when beginning a
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-coding-scope"
  target="_blank" rel="noopener">video coding scope</a>.

- `stdSyntaxFlags` is a bitmask of
  [VkVideoEncodeH264StdFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264StdFlagBitsKHR.html)
  indicating capabilities related to H.264 syntax elements.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkVideoEncodeH264CapabilitiesKHR-sType-sType"
  id="VUID-VkVideoEncodeH264CapabilitiesKHR-sType-sType"></a>
  VUID-VkVideoEncodeH264CapabilitiesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_CAPABILITIES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_h264](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_h264.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoEncodeH264CapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilityFlagsKHR.html),
[VkVideoEncodeH264StdFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264StdFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeH264CapabilitiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
