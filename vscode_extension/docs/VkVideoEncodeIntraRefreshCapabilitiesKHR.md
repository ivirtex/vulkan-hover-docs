# VkVideoEncodeIntraRefreshCapabilitiesKHR(3) Manual Page

## Name

VkVideoEncodeIntraRefreshCapabilitiesKHR - Structure describing video encode intra refresh capabilities for a video profile



## [](#_c_specification)C Specification

When calling [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) with `pVideoProfile->videoCodecOperation` specifying an encode operation, the [VkVideoEncodeIntraRefreshCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeIntraRefreshCapabilitiesKHR.html) structure **can** be included in the `pNext` chain of the [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html) structure to retrieve capabilities specific to video encode intra refresh.

The `VkVideoEncodeIntraRefreshCapabilitiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_intra_refresh
typedef struct VkVideoEncodeIntraRefreshCapabilitiesKHR {
    VkStructureType                          sType;
    void*                                    pNext;
    VkVideoEncodeIntraRefreshModeFlagsKHR    intraRefreshModes;
    uint32_t                                 maxIntraRefreshCycleDuration;
    uint32_t                                 maxIntraRefreshActiveReferencePictures;
    VkBool32                                 partitionIndependentIntraRefreshRegions;
    VkBool32                                 nonRectangularIntraRefreshRegions;
} VkVideoEncodeIntraRefreshCapabilitiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `intraRefreshModes` is a bitmask of [VkVideoEncodeIntraRefreshModeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeIntraRefreshModeFlagBitsKHR.html) values indicating the set of supported [intra refresh modes](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-intra-refresh-modes).
- `maxIntraRefreshCycleDuration` specifies the maximum supported [intra refresh cycle duration](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-intra-refresh-cycle-duration).
- `maxIntraRefreshActiveReferencePictures` is the maximum number of [active reference pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#active-reference-pictures) when encoding pictures with [intra refresh](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-intra-refresh) enabled. This capability indicates additional restrictions beyond the maximum number of [active reference pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#active-reference-pictures) supported by the video profile, as reported in [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html)::`maxActiveReferencePictures` and the maximum requested at video session creation time in [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionCreateInfoKHR.html)::`maxActiveReferencePictures`.
- `partitionIndependentIntraRefreshRegions` specifies whether the implementation supports intra refresh regions that are independent of the picture partitioning used during encoding. If it is `VK_TRUE`, then pictures **can** be encoded with multiple picture partitions, independent of the used intra refresh mode. Otherwise, pictures **cannot** be encoded with multiple picture partitions with any intra refresh mode other than `VK_VIDEO_ENCODE_INTRA_REFRESH_MODE_PER_PICTURE_PARTITION_BIT_KHR`.
  
  Note
  
  This capability is only indicative for [AV1 encode profiles](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-profile) and does not impose any restrictions on the application as implementations may change the application requested picture partitioning according to implementation-specific restrictions.
- `nonRectangularIntraRefreshRegions` specifies whether the implementation supports non-rectangular intra refresh regions.
  
  Note
  
  If this capability is not supported, then using per picture partition intra refresh may impose additional restrictions on the number of picture partitions a picture can be encoded with.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeIntraRefreshCapabilitiesKHR-sType-sType)VUID-VkVideoEncodeIntraRefreshCapabilitiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_INTRA_REFRESH_CAPABILITIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_intra\_refresh](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_intra_refresh.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoEncodeIntraRefreshModeFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeIntraRefreshModeFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeIntraRefreshCapabilitiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0