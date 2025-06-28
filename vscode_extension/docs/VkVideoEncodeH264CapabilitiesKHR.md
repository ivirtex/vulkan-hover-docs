# VkVideoEncodeH264CapabilitiesKHR(3) Manual Page

## Name

VkVideoEncodeH264CapabilitiesKHR - Structure describing H.264 encode capabilities



## [](#_c_specification)C Specification

When calling [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) to query the capabilities for an [H.264 encode profile](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-profile), the [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html)::`pNext` chain **must** include a `VkVideoEncodeH264CapabilitiesKHR` structure that will be filled with the profile-specific capabilities.

The `VkVideoEncodeH264CapabilitiesKHR` structure is defined as:

```c++
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

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkVideoEncodeH264CapabilityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264CapabilityFlagBitsKHR.html) indicating supported H.264 encoding capabilities.
- `maxLevelIdc` is a `StdVideoH264LevelIdc` value indicating the maximum H.264 level supported by the profile, where enum constant `STD_VIDEO_H264_LEVEL_IDC_<major>_<minor>` identifies H.264 level `<major>.<minor>` as defined in section A.3 of the [ITU-T H.264 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h264).
- `maxSliceCount` indicates the maximum number of slices that **can** be encoded for a single picture. Further restrictions **may** apply to the number of slices that **can** be encoded for a single picture depending on other capabilities and codec-specific rules.
- `maxPPictureL0ReferenceCount` indicates the maximum number of reference pictures the implementation supports in the reference list L0 for [P pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-p-pic).
  
  Note
  
  As implementations **may** [override](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-overrides) the reference lists, `maxPPictureL0ReferenceCount` does not limit the number of elements that the application **can** specify in the L0 reference list for P pictures. However, if `maxPPictureL0ReferenceCount` is zero, then the use of P pictures is not allowed.
- `maxBPictureL0ReferenceCount` indicates the maximum number of reference pictures the implementation supports in the reference list L0 for [B pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-b-pic).
- `maxL1ReferenceCount` indicates the maximum number of reference pictures the implementation supports in the reference list L1 if encoding of [B pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-b-pic) is supported.
  
  Note
  
  As implementations **may** [override](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-overrides) the reference lists, `maxBPictureL0ReferenceCount` and `maxL1ReferenceCount` does not limit the number of elements that the application **can** specify in the L0 and L1 reference lists for B pictures. However, if `maxBPictureL0ReferenceCount` and `maxL1ReferenceCount` are both zero, then the use of B pictures is not allowed.
- `maxTemporalLayerCount` indicates the maximum number of H.264 temporal layers supported by the implementation.
- `expectDyadicTemporalLayerPattern` indicates that the implementation’s rate control algorithms expect the application to use a [dyadic temporal layer pattern](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-layer-pattern-dyadic) when encoding multiple temporal layers.
- `minQp` indicates the minimum QP value supported.
- `maxQp` indicates the maximum QP value supported.
- `prefersGopRemainingFrames` indicates that the implementation’s rate control algorithm prefers the application to specify the number of frames of each type [remaining](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-gop-remaining-frames) in the current [group of pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-gop) when beginning a [video coding scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-coding-scope).
- `requiresGopRemainingFrames` indicates that the implementation’s rate control algorithm requires the application to specify the number of frames of each type [remaining](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-gop-remaining-frames) in the current [group of pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-gop) when beginning a [video coding scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-coding-scope).
- `stdSyntaxFlags` is a bitmask of [VkVideoEncodeH264StdFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264StdFlagBitsKHR.html) indicating capabilities related to H.264 syntax elements.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeH264CapabilitiesKHR-sType-sType)VUID-VkVideoEncodeH264CapabilitiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_CAPABILITIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_h264](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h264.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoEncodeH264CapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264CapabilityFlagsKHR.html), [VkVideoEncodeH264StdFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264StdFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeH264CapabilitiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0