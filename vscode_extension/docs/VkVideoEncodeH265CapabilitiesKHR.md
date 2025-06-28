# VkVideoEncodeH265CapabilitiesKHR(3) Manual Page

## Name

VkVideoEncodeH265CapabilitiesKHR - Structure describing H.265 encode capabilities



## [](#_c_specification)C Specification

When calling [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) to query the capabilities for an [H.265 encode profile](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-profile), the [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html)::`pNext` chain **must** include a `VkVideoEncodeH265CapabilitiesKHR` structure that will be filled with the profile-specific capabilities.

The `VkVideoEncodeH265CapabilitiesKHR` structure is defined as:

```c++
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

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkVideoEncodeH265CapabilityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CapabilityFlagBitsKHR.html) indicating supported H.265 encoding capabilities.
- `maxLevelIdc` is a `StdVideoH265LevelIdc` value indicating the maximum H.265 level supported by the profile, where enum constant `STD_VIDEO_H265_LEVEL_IDC_<major>_<minor>` identifies H.265 level `<major>.<minor>` as defined in section A.4 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265).
- `maxSliceSegmentCount` indicates the maximum number of slice segments that **can** be encoded for a single picture. Further restrictions **may** apply to the number of slice segments that **can** be encoded for a single picture depending on other capabilities and codec-specific rules.
- `maxTiles` indicates the maximum number of H.265 tile columns and rows, as defined in sections 3.175 and 3.176 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265) that **can** be encoded for a single picture. Further restrictions **may** apply to the number of H.265 tiles that **can** be encoded for a single picture depending on other capabilities and codec-specific rules.
- `ctbSizes` is a bitmask of [VkVideoEncodeH265CtbSizeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CtbSizeFlagBitsKHR.html) describing the supported CTB sizes.
- `transformBlockSizes` is a bitmask of [VkVideoEncodeH265TransformBlockSizeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265TransformBlockSizeFlagBitsKHR.html) describing the supported transform block sizes.
- `maxPPictureL0ReferenceCount` indicates the maximum number of reference pictures the implementation supports in the reference list L0 for [P pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-p-pic).
  
  Note
  
  As implementations **may** [override](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-overrides) the reference lists, `maxPPictureL0ReferenceCount` does not limit the number of elements that the application **can** specify in the L0 reference list for P pictures. However, if `maxPPictureL0ReferenceCount` is zero, then the use of P pictures is not allowed. In case of H.265 encoding, pictures **can** be encoded using only forward prediction even if P pictures are not supported, as the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265) supports *generalized P &amp; B frames* (also known as low delay B frames) whereas B frames **can** refer to past frames through both the L0 and L1 reference lists.
- `maxBPictureL0ReferenceCount` indicates the maximum number of reference pictures the implementation supports in the reference list L0 for [B pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-b-pic).
- `maxL1ReferenceCount` indicates the maximum number of reference pictures the implementation supports in the reference list L1 if encoding of [B pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-b-pic) is supported.
  
  Note
  
  As implementations **may** [override](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-overrides) the reference lists, `maxBPictureL0ReferenceCount` and `maxL1ReferenceCount` does not limit the number of elements that the application **can** specify in the L0 and L1 reference lists for B pictures. However, if `maxBPictureL0ReferenceCount` and `maxL1ReferenceCount` are both zero, then the use of B pictures is not allowed.
- `maxSubLayerCount` indicates the maximum number of H.265 sub-layers supported by the implementation.
- `expectDyadicTemporalSubLayerPattern` indicates that the implementation’s rate control algorithms expect the application to use a [dyadic temporal sub-layer pattern](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-layer-pattern-dyadic) when encoding multiple temporal sub-layers.
- `minQp` indicates the minimum QP value supported.
- `maxQp` indicates the maximum QP value supported.
- `prefersGopRemainingFrames` indicates that the implementation’s rate control algorithm prefers the application to specify the number of frames of each type [remaining](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-gop-remaining-frames) in the current [group of pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-gop) when beginning a [video coding scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-coding-scope).
- `requiresGopRemainingFrames` indicates that the implementation’s rate control algorithm requires the application to specify the number of frames of each type [remaining](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-gop-remaining-frames) in the current [group of pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-gop) when beginning a [video coding scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-coding-scope).
- `stdSyntaxFlags` is a bitmask of [VkVideoEncodeH265StdFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265StdFlagBitsKHR.html) indicating capabilities related to H.265 syntax elements.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeH265CapabilitiesKHR-sType-sType)VUID-VkVideoEncodeH265CapabilitiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_CAPABILITIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_h265](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h265.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoEncodeH265CapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CapabilityFlagsKHR.html), [VkVideoEncodeH265CtbSizeFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CtbSizeFlagsKHR.html), [VkVideoEncodeH265StdFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265StdFlagsKHR.html), [VkVideoEncodeH265TransformBlockSizeFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265TransformBlockSizeFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeH265CapabilitiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0