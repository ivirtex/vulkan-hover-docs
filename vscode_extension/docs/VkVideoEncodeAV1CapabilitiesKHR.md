# VkVideoEncodeAV1CapabilitiesKHR(3) Manual Page

## Name

VkVideoEncodeAV1CapabilitiesKHR - Structure describing AV1 encode capabilities



## [](#_c_specification)C Specification

When calling [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) to query the capabilities for an [AV1 encode profile](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-profile), the [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html)::`pNext` chain **must** include a `VkVideoEncodeAV1CapabilitiesKHR` structure that will be filled with the profile-specific capabilities.

The `VkVideoEncodeAV1CapabilitiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_av1
typedef struct VkVideoEncodeAV1CapabilitiesKHR {
    VkStructureType                           sType;
    void*                                     pNext;
    VkVideoEncodeAV1CapabilityFlagsKHR        flags;
    StdVideoAV1Level                          maxLevel;
    VkExtent2D                                codedPictureAlignment;
    VkExtent2D                                maxTiles;
    VkExtent2D                                minTileSize;
    VkExtent2D                                maxTileSize;
    VkVideoEncodeAV1SuperblockSizeFlagsKHR    superblockSizes;
    uint32_t                                  maxSingleReferenceCount;
    uint32_t                                  singleReferenceNameMask;
    uint32_t                                  maxUnidirectionalCompoundReferenceCount;
    uint32_t                                  maxUnidirectionalCompoundGroup1ReferenceCount;
    uint32_t                                  unidirectionalCompoundReferenceNameMask;
    uint32_t                                  maxBidirectionalCompoundReferenceCount;
    uint32_t                                  maxBidirectionalCompoundGroup1ReferenceCount;
    uint32_t                                  maxBidirectionalCompoundGroup2ReferenceCount;
    uint32_t                                  bidirectionalCompoundReferenceNameMask;
    uint32_t                                  maxTemporalLayerCount;
    uint32_t                                  maxSpatialLayerCount;
    uint32_t                                  maxOperatingPoints;
    uint32_t                                  minQIndex;
    uint32_t                                  maxQIndex;
    VkBool32                                  prefersGopRemainingFrames;
    VkBool32                                  requiresGopRemainingFrames;
    VkVideoEncodeAV1StdFlagsKHR               stdSyntaxFlags;
} VkVideoEncodeAV1CapabilitiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkVideoEncodeAV1CapabilityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1CapabilityFlagBitsKHR.html) indicating supported AV1 encoding capabilities.
- `maxLevel` is a `StdVideoAV1Level` value indicating the maximum AV1 level supported by the profile, as defined in section A.3 of the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1).
- `codedPictureAlignment` indicates the alignment at which the implementation will code pictures. This capability does not impose any valid usage constraints on the application. However, depending on the `codedExtent` of the encode input picture resource, this capability **may** result in a change of the resolution of the encoded picture, as described in more detail below.
- `maxTiles` indicates the maximum number of AV1 tile columns and rows the implementation supports.
- `minTileSize` indicates the minimum extent of individual AV1 tiles the implementation supports.
- `maxTileSize` indicates the maximum extent of individual AV1 tiles the implementation supports.
- `superblockSizes` is a bitmask of [VkVideoEncodeAV1SuperblockSizeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1SuperblockSizeFlagBitsKHR.html) values indicating the supported AV1 superblock sizes.
- `maxSingleReferenceCount` indicates the maximum number of reference pictures the implementation supports when using [single reference prediction mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-prediction-modes).
- `singleReferenceNameMask` is a bitmask of supported [AV1 reference names](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-reference-names) when using [single reference prediction mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-prediction-modes).
- `maxUnidirectionalCompoundReferenceCount` indicates the maximum number of reference pictures the implementation supports when using [unidirectional compound prediction mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-prediction-modes).
- `maxUnidirectionalCompoundGroup1ReferenceCount` indicates the maximum number of reference pictures the implementation supports when using [unidirectional compound prediction mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-prediction-modes) from reference frame group 1, as defined in section 6.10.24 of the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1).
- `unidirectionalCompoundReferenceNameMask` is a bitmask of supported [AV1 reference names](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-reference-names) when using [unidirectional compound prediction mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-prediction-modes).
- `maxBidirectionalCompoundReferenceCount` indicates the maximum number of reference pictures the implementation supports when using [bidirectional compound prediction mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-prediction-modes).
- `maxBidirectionalCompoundGroup1ReferenceCount` indicates the maximum number of reference pictures the implementation supports when using [bidirectional compound prediction mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-prediction-modes) from reference frame group 1, as defined in section 6.10.24 of the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1).
- `maxBidirectionalCompoundGroup2ReferenceCount` indicates the maximum number of reference pictures the implementation supports when using [bidirectional compound prediction mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-prediction-modes) from reference frame group 2, as defined in section 6.10.24 of the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1).
- `bidirectionalCompoundReferenceNameMask` is a bitmask of supported [AV1 reference names](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-reference-names) when using [bidirectional compound prediction mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-prediction-modes).
- `maxTemporalLayerCount` indicates the maximum number of AV1 temporal layers supported by the implementation.
- `maxSpatialLayerCount` indicates the maximum number of AV1 spatial layers supported by the implementation.
- `maxOperatingPoints` indicates the maximum number of AV1 operating points supported by the implementation.
- `minQIndex` indicates the minimum quantizer index value supported.
- `maxQIndex` indicates the maximum quantizer index value supported.
- `prefersGopRemainingFrames` indicates that the implementation’s rate control algorithm prefers the application to specify the number of frames in each [AV1 rate control group](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-rate-control-group) [remaining](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-gop-remaining-frames) in the current [group of pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-gop) when beginning a [video coding scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-coding-scope).
- `requiresGopRemainingFrames` indicates that the implementation’s rate control algorithm requires the application to specify the number of frames in each [AV1 rate control group](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-rate-control-group) [remaining](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-gop-remaining-frames) in the current [group of pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-gop) when beginning a [video coding scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-coding-scope).
- `stdSyntaxFlags` is a bitmask of [VkVideoEncodeAV1StdFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1StdFlagBitsKHR.html) indicating capabilities related to AV1 syntax elements.

## [](#_description)Description

`singleReferenceNameMask`, `unidirectionalCompoundReferenceNameMask`, and `bidirectionalCompoundReferenceNameMask` are encoded such that when bit index i is set, it indicates support for the [AV1 reference name](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-reference-names) `STD_VIDEO_AV1_REFERENCE_NAME_LAST_FRAME` + i.

Note

These masks indicate which elements of the `referenceNameSlotIndices` member of [VkVideoEncodeAV1PictureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1PictureInfoKHR.html) are supported to be used by the implementation. It is important to note that both the bits of these masks and the elements of `referenceNameSlotIndices` are indexed such that the first value specifies the support bit and DPB slot index, respectively, for the AV1 reference name `STD_VIDEO_AV1_REFERENCE_NAME_LAST_FRAME` (i.e. there is no bit or element for `STD_VIDEO_AV1_REFERENCE_NAME_INTRA_FRAME`).

`codedPictureAlignment` provides information about implementation limitations to encode arbitrary resolutions. In particular, some implementations **may** not be able to generate bitstreams aligned to the requirements of the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1) (8x8). In such cases, the implementation **may** [override the width and height of the bitstream](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-resolution-override), in order to produce a bitstream compliant to the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1). If such an override occurs, the encoded resolution of the coded picture is enlargened, with the texel values used for the texel coordinates outside of the bounds of the `codedExtent` of the encode input picture resource being first governed by the rules regarding the [encode input picture granularity](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-input-picture-granularity). Any texel values outside of the region described by the encode input picture granularity are implementation-defined. Implementations **should** use well-defined values to minimize impact on the produced encoded content.

Note

This capability does not impose additional application requirements. However, these overrides change the effective resolution of the bitstream and add padding pixels. Applications sensitive to such overrides **can** use this capability and the corresponding [override behavior](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-resolution-override) to compute the cropping needed to reproduce the original input of the encoding and transmit it in a side channel (i.e. by using cropping fields available in a container). Additionally, applications **can** explicitly consider this alignment in their coded extent, to avoid implementation-defined texel values being included in the encoded content.

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeAV1CapabilitiesKHR-sType-sType)VUID-VkVideoEncodeAV1CapabilitiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_AV1_CAPABILITIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_av1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_av1.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoEncodeAV1CapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1CapabilityFlagsKHR.html), [VkVideoEncodeAV1StdFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1StdFlagsKHR.html), [VkVideoEncodeAV1SuperblockSizeFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1SuperblockSizeFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeAV1CapabilitiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0