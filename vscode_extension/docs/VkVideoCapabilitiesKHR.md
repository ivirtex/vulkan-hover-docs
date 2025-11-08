# VkVideoCapabilitiesKHR(3) Manual Page

## Name

VkVideoCapabilitiesKHR - Structure describing general video capabilities for a video profile



## [](#_c_specification)C Specification

The `VkVideoCapabilitiesKHR` structure is defined as:

```c++
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

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkVideoCapabilityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilityFlagBitsKHR.html) specifying capability flags.
- `minBitstreamBufferOffsetAlignment` is the minimum alignment for bitstream buffer offsets.
- `minBitstreamBufferSizeAlignment` is the minimum alignment for bitstream buffer range sizes.
- `pictureAccessGranularity` is the granularity at which image access to video picture resources happen.
- `minCodedExtent` is the minimum width and height of the coded frames.
- `maxCodedExtent` is the maximum width and height of the coded frames.
- `maxDpbSlots` is the maximum number of [DPB slots](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot) supported by a single video session.
- `maxActiveReferencePictures` is the maximum number of [active reference pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#active-reference-pictures) a single video coding operation **can** use.
- []()`stdHeaderVersion` is a [VkExtensionProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtensionProperties.html) structure reporting the Video Std header name and version supported for the video profile.

## [](#_description)Description

Note

It is common for video compression standards to allow using all reference pictures associated with active DPB slots as active reference pictures, hence for video decode profiles the values returned in `maxDpbSlots` and `maxActiveReferencePictures` are often equal. Similarly, in case of video decode profiles supporting field pictures the value of `maxActiveReferencePictures` often equals `maxDpbSlots` × 2.

Valid Usage (Implicit)

- [](#VUID-VkVideoCapabilitiesKHR-sType-sType)VUID-VkVideoCapabilitiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_CAPABILITIES_KHR`
- [](#VUID-VkVideoCapabilitiesKHR-pNext-pNext)VUID-VkVideoCapabilitiesKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkVideoDecodeAV1CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeAV1CapabilitiesKHR.html), [VkVideoDecodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeCapabilitiesKHR.html), [VkVideoDecodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264CapabilitiesKHR.html), [VkVideoDecodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265CapabilitiesKHR.html), [VkVideoDecodeVP9CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeVP9CapabilitiesKHR.html), [VkVideoEncodeAV1CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1CapabilitiesKHR.html), [VkVideoEncodeAV1QuantizationMapCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1QuantizationMapCapabilitiesKHR.html), [VkVideoEncodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeCapabilitiesKHR.html), [VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264CapabilitiesKHR.html), [VkVideoEncodeH264QuantizationMapCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264QuantizationMapCapabilitiesKHR.html), [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CapabilitiesKHR.html), [VkVideoEncodeH265QuantizationMapCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265QuantizationMapCapabilitiesKHR.html), [VkVideoEncodeIntraRefreshCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeIntraRefreshCapabilitiesKHR.html), [VkVideoEncodeQuantizationMapCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQuantizationMapCapabilitiesKHR.html), or [VkVideoEncodeRgbConversionCapabilitiesVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbConversionCapabilitiesVALVE.html)
- [](#VUID-VkVideoCapabilitiesKHR-sType-unique)VUID-VkVideoCapabilitiesKHR-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkExtensionProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtensionProperties.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoCapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilityFlagsKHR.html), [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoCapabilitiesKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0