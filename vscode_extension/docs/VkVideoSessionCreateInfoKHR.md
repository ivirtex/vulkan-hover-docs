# VkVideoSessionCreateInfoKHR(3) Manual Page

## Name

VkVideoSessionCreateInfoKHR - Structure specifying parameters of a newly created video session



## [](#_c_specification)C Specification

The [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionCreateInfoKHR.html) structure is defined as:

```c++
// Provided by VK_KHR_video_queue
typedef struct VkVideoSessionCreateInfoKHR {
    VkStructureType                 sType;
    const void*                     pNext;
    uint32_t                        queueFamilyIndex;
    VkVideoSessionCreateFlagsKHR    flags;
    const VkVideoProfileInfoKHR*    pVideoProfile;
    VkFormat                        pictureFormat;
    VkExtent2D                      maxCodedExtent;
    VkFormat                        referencePictureFormat;
    uint32_t                        maxDpbSlots;
    uint32_t                        maxActiveReferencePictures;
    const VkExtensionProperties*    pStdHeaderVersion;
} VkVideoSessionCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `queueFamilyIndex` is the index of the queue family the created video session will be used with.
- `flags` is a bitmask of [VkVideoSessionCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionCreateFlagBitsKHR.html) specifying creation flags.
- `pVideoProfile` is a pointer to a [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html) structure specifying the video profile the created video session will be used with.
- `pictureFormat` is the image format the created video session will be used with. If `pVideoProfile->videoCodecOperation` specifies a decode operation, then `pictureFormat` is the image format of [decode output pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-output-picture) usable with the created video session. If `pVideoProfile->videoCodecOperation` specifies an encode operation, then `pictureFormat` is the image format of [encode input pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-input-picture) usable with the created video session.
- `maxCodedExtent` is the maximum width and height of the coded frames the created video session will be used with.
- `referencePictureFormat` is the image format of [reference pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#reference-picture) stored in the [DPB](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb) the created video session will be used with.
- `maxDpbSlots` is the maximum number of [DPB Slots](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot) that **can** be used with the created video session.
- `maxActiveReferencePictures` is the maximum number of [active reference pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#active-reference-pictures) that **can** be used in a single video coding operation using the created video session.
- `pStdHeaderVersion` is a pointer to a [VkExtensionProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtensionProperties.html) structure requesting the Video Std header version to use for the `videoCodecOperation` specified in `pVideoProfile`.

## [](#_description)Description

Valid Usage

- [](#VUID-VkVideoSessionCreateInfoKHR-protectedMemory-07189)VUID-VkVideoSessionCreateInfoKHR-protectedMemory-07189  
  If the [`protectedMemory`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-protectedMemory) feature is not enabled or if [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html)::`flags` does not include `VK_VIDEO_CAPABILITY_PROTECTED_CONTENT_BIT_KHR`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the video profile specified by `pVideoProfile`, then `flags` **must** not include `VK_VIDEO_SESSION_CREATE_PROTECTED_CONTENT_BIT_KHR`
- [](#VUID-VkVideoSessionCreateInfoKHR-flags-08371)VUID-VkVideoSessionCreateInfoKHR-flags-08371  
  If `flags` includes `VK_VIDEO_SESSION_CREATE_INLINE_QUERIES_BIT_KHR`, then [`videoMaintenance1`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-videoMaintenance1) **must** be enabled
- [](#VUID-VkVideoSessionCreateInfoKHR-flags-10398)VUID-VkVideoSessionCreateInfoKHR-flags-10398  
  If `flags` includes `VK_VIDEO_SESSION_CREATE_INLINE_SESSION_PARAMETERS_BIT_KHR`, then [`videoMaintenance2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-videoMaintenance2) **must** be enabled
- [](#VUID-VkVideoSessionCreateInfoKHR-flags-10399)VUID-VkVideoSessionCreateInfoKHR-flags-10399  
  If `flags` includes `VK_VIDEO_SESSION_CREATE_INLINE_SESSION_PARAMETERS_BIT_KHR`, then `pVideoProfile->videoCodecOperation` **must** specify a decode operation
- [](#VUID-VkVideoSessionCreateInfoKHR-flags-10264)VUID-VkVideoSessionCreateInfoKHR-flags-10264  
  If `flags` includes `VK_VIDEO_SESSION_CREATE_ALLOW_ENCODE_QUANTIZATION_DELTA_MAP_BIT_KHR` or `VK_VIDEO_SESSION_CREATE_ALLOW_ENCODE_EMPHASIS_MAP_BIT_KHR`, then the [`videoEncodeQuantizationMap`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-videoEncodeQuantizationMap) feature **must** be enabled
- [](#VUID-VkVideoSessionCreateInfoKHR-flags-10265)VUID-VkVideoSessionCreateInfoKHR-flags-10265  
  If `flags` includes `VK_VIDEO_SESSION_CREATE_ALLOW_ENCODE_QUANTIZATION_DELTA_MAP_BIT_KHR` or `VK_VIDEO_SESSION_CREATE_ALLOW_ENCODE_EMPHASIS_MAP_BIT_KHR`, then `pVideoProfile->videoCodecOperation` **must** specify an encode operation
- [](#VUID-VkVideoSessionCreateInfoKHR-flags-10266)VUID-VkVideoSessionCreateInfoKHR-flags-10266  
  If `flags` includes `VK_VIDEO_SESSION_CREATE_ALLOW_ENCODE_QUANTIZATION_DELTA_MAP_BIT_KHR`, then it **must** not also include `VK_VIDEO_SESSION_CREATE_ALLOW_ENCODE_EMPHASIS_MAP_BIT_KHR`
- [](#VUID-VkVideoSessionCreateInfoKHR-flags-10267)VUID-VkVideoSessionCreateInfoKHR-flags-10267  
  If [VkVideoEncodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeCapabilitiesKHR.html)::`flags` does not include `VK_VIDEO_ENCODE_CAPABILITY_QUANTIZATION_DELTA_MAP_BIT_KHR`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the video profile specified by `pVideoProfile`, then `flags` **must** not include `VK_VIDEO_SESSION_CREATE_ALLOW_ENCODE_QUANTIZATION_DELTA_MAP_BIT_KHR`
- [](#VUID-VkVideoSessionCreateInfoKHR-flags-10268)VUID-VkVideoSessionCreateInfoKHR-flags-10268  
  If [VkVideoEncodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeCapabilitiesKHR.html)::`flags` does not include `VK_VIDEO_ENCODE_CAPABILITY_EMPHASIS_MAP_BIT_KHR`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the video profile specified by `pVideoProfile`, then `flags` **must** not include `VK_VIDEO_SESSION_CREATE_ALLOW_ENCODE_EMPHASIS_MAP_BIT_KHR`
- [](#VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-04845)VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-04845  
  `pVideoProfile` **must** be a [supported video profile](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-profile-support)
- [](#VUID-VkVideoSessionCreateInfoKHR-maxDpbSlots-04847)VUID-VkVideoSessionCreateInfoKHR-maxDpbSlots-04847  
  `maxDpbSlots` **must** be less than or equal to [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html)::`maxDpbSlots`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the video profile specified by `pVideoProfile`
- [](#VUID-VkVideoSessionCreateInfoKHR-maxActiveReferencePictures-04849)VUID-VkVideoSessionCreateInfoKHR-maxActiveReferencePictures-04849  
  `maxActiveReferencePictures` **must** be less than or equal to [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html)::`maxActiveReferencePictures`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the video profile specified by `pVideoProfile`
- [](#VUID-VkVideoSessionCreateInfoKHR-maxDpbSlots-04850)VUID-VkVideoSessionCreateInfoKHR-maxDpbSlots-04850  
  If either `maxDpbSlots` or `maxActiveReferencePictures` is `0`, then both **must** be `0`
- [](#VUID-VkVideoSessionCreateInfoKHR-maxCodedExtent-04851)VUID-VkVideoSessionCreateInfoKHR-maxCodedExtent-04851  
  `maxCodedExtent` **must** be between [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html)::`minCodedExtent` and [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html)::`maxCodedExtent`, inclusive, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the video profile specified by `pVideoProfile`
- [](#VUID-VkVideoSessionCreateInfoKHR-referencePictureFormat-04852)VUID-VkVideoSessionCreateInfoKHR-referencePictureFormat-04852  
  If `pVideoProfile->videoCodecOperation` specifies a decode operation and `maxActiveReferencePictures` is greater than `0`, then `referencePictureFormat` **must** be one of the supported decode DPB formats, as returned by [vkGetPhysicalDeviceVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoFormatPropertiesKHR.html) in [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatPropertiesKHR.html)::`format` when called with the `imageUsage` member of its `pVideoFormatInfo` parameter containing `VK_IMAGE_USAGE_VIDEO_DECODE_DPB_BIT_KHR`, and with a [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileListInfoKHR.html) structure specified in the `pNext` chain of its `pVideoFormatInfo` parameter whose `pProfiles` member contains an element matching `pVideoProfile`
- [](#VUID-VkVideoSessionCreateInfoKHR-referencePictureFormat-06814)VUID-VkVideoSessionCreateInfoKHR-referencePictureFormat-06814  
  If `pVideoProfile->videoCodecOperation` specifies an encode operation and `maxActiveReferencePictures` is greater than `0`, then `referencePictureFormat` **must** be one of the supported decode DPB formats, as returned by then `referencePictureFormat` **must** be one of the supported encode DPB formats, as returned by [vkGetPhysicalDeviceVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoFormatPropertiesKHR.html) in [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatPropertiesKHR.html)::`format` when called with the `imageUsage` member of its `pVideoFormatInfo` parameter containing `VK_IMAGE_USAGE_VIDEO_ENCODE_DPB_BIT_KHR`, and with a [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileListInfoKHR.html) structure specified in the `pNext` chain of its `pVideoFormatInfo` parameter whose `pProfiles` member contains an element matching `pVideoProfile`
- [](#VUID-VkVideoSessionCreateInfoKHR-pictureFormat-04853)VUID-VkVideoSessionCreateInfoKHR-pictureFormat-04853  
  If `pVideoProfile->videoCodecOperation` specifies a decode operation, then `pictureFormat` **must** be one of the supported decode output formats, as returned by [vkGetPhysicalDeviceVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoFormatPropertiesKHR.html) in [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatPropertiesKHR.html)::`format` when called with the `imageUsage` member of its `pVideoFormatInfo` parameter containing `VK_IMAGE_USAGE_VIDEO_DECODE_DST_BIT_KHR`, and with a [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileListInfoKHR.html) structure specified in the `pNext` chain of its `pVideoFormatInfo` parameter whose `pProfiles` member contains an element matching `pVideoProfile`
- [](#VUID-VkVideoSessionCreateInfoKHR-pictureFormat-04854)VUID-VkVideoSessionCreateInfoKHR-pictureFormat-04854  
  If `pVideoProfile->videoCodecOperation` specifies an encode operation, then `pictureFormat` **must** be one of the supported encode input formats, as returned by [vkGetPhysicalDeviceVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoFormatPropertiesKHR.html) in [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatPropertiesKHR.html)::`format` when called with the `imageUsage` member of its `pVideoFormatInfo` parameter containing `VK_IMAGE_USAGE_VIDEO_ENCODE_SRC_BIT_KHR`, and with a [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileListInfoKHR.html) structure specified in the `pNext` chain of its `pVideoFormatInfo` parameter whose `pProfiles` member contains an element matching `pVideoProfile`
- [](#VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-10835)VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-10835  
  If `pVideoProfile->videoCodecOperation` specifies an encode operation, the `pNext` chain of this structure includes a [VkVideoEncodeSessionIntraRefreshCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeSessionIntraRefreshCreateInfoKHR.html) structure, and its `intraRefreshMode` member is not `VK_VIDEO_ENCODE_INTRA_REFRESH_MODE_NONE_KHR`, then the [`videoEncodeIntraRefresh`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-videoEncodeIntraRefresh) feature **must** be enabled
- [](#VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-10836)VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-10836  
  If `pVideoProfile->videoCodecOperation` specifies an encode operation, the `pNext` chain of this structure includes a [VkVideoEncodeSessionIntraRefreshCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeSessionIntraRefreshCreateInfoKHR.html) structure, and its `intraRefreshMode` member is not `VK_VIDEO_ENCODE_INTRA_REFRESH_MODE_NONE_KHR`, then `intraRefreshMode` **must** specify one of the bits included in [VkVideoEncodeIntraRefreshCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeIntraRefreshCapabilitiesKHR.html)::`intraRefreshModes`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the video profile specified by `pVideoProfile`
- [](#VUID-VkVideoSessionCreateInfoKHR-pStdHeaderVersion-07190)VUID-VkVideoSessionCreateInfoKHR-pStdHeaderVersion-07190  
  `pStdHeaderVersion->extensionName` **must** match [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html)::`stdHeaderVersion.extensionName`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the video profile specified by `pVideoProfile`
- [](#VUID-VkVideoSessionCreateInfoKHR-pStdHeaderVersion-07191)VUID-VkVideoSessionCreateInfoKHR-pStdHeaderVersion-07191  
  `pStdHeaderVersion->specVersion` **must** be less than or equal to [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html)::`stdHeaderVersion.specVersion`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the video profile specified by `pVideoProfile`
- [](#VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-10793)VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-10793  
  If `pVideoProfile->videoCodecOperation` is `VK_VIDEO_CODEC_OPERATION_DECODE_VP9_BIT_KHR`, then the [`videoDecodeVP9`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-videoDecodeVP9) feature **must** be enabled
- [](#VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-08251)VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-08251  
  If `pVideoProfile->videoCodecOperation` is `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR` and the `pNext` chain of this structure includes a [VkVideoEncodeH264SessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionCreateInfoKHR.html) structure, then its `maxLevelIdc` member **must** be less than or equal to [VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264CapabilitiesKHR.html)::`maxLevelIdc`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the video profile specified in `pVideoProfile`
- [](#VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-08252)VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-08252  
  If `pVideoProfile->videoCodecOperation` is `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR` and the `pNext` chain of this structure includes a [VkVideoEncodeH265SessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionCreateInfoKHR.html) structure, then its `maxLevelIdc` member **must** be less than or equal to [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`maxLevelIdc`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the video profile specified in `pVideoProfile`
- [](#VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-10269)VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-10269  
  If `pVideoProfile->videoCodecOperation` is `VK_VIDEO_CODEC_OPERATION_ENCODE_AV1_BIT_KHR`, then the [`videoEncodeAV1`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-videoEncodeAV1) feature **must** be enabled
- [](#VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-10270)VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-10270  
  If `pVideoProfile->videoCodecOperation` is `VK_VIDEO_CODEC_OPERATION_ENCODE_AV1_BIT_KHR` and the `pNext` chain of this structure includes a [VkVideoEncodeAV1SessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1SessionCreateInfoKHR.html) structure, then its `maxLevel` member **must** be less than or equal to [VkVideoEncodeAV1CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1CapabilitiesKHR.html)::`maxLevel`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the video profile specified in `pVideoProfile`
- [](#VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-10923)VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-10923  
  If the `pVideoProfile->pNext` chain includes a [VkVideoEncodeProfileRgbConversionInfoVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeProfileRgbConversionInfoVALVE.html) structure, then the [`videoEncodeRgbConversion`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-videoEncodeRgbConversion) feature **must** be enabled
- [](#VUID-VkVideoSessionCreateInfoKHR-pNext-10924)VUID-VkVideoSessionCreateInfoKHR-pNext-10924  
  If a [VkVideoEncodeProfileRgbConversionInfoVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeProfileRgbConversionInfoVALVE.html) structure is included in the `pNext` chain of `pVideoProfile` and `VkVideoEncodeProfileRgbConversionInfoVALVE`::`performEncodeRgbConversion` is enabled, a [VkVideoEncodeSessionRgbConversionCreateInfoVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeSessionRgbConversionCreateInfoVALVE.html) structure **must** be included in the `pNext` chain of [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionCreateInfoKHR.html).

Valid Usage (Implicit)

- [](#VUID-VkVideoSessionCreateInfoKHR-sType-sType)VUID-VkVideoSessionCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_SESSION_CREATE_INFO_KHR`
- [](#VUID-VkVideoSessionCreateInfoKHR-pNext-pNext)VUID-VkVideoSessionCreateInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkVideoEncodeAV1SessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1SessionCreateInfoKHR.html), [VkVideoEncodeH264SessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionCreateInfoKHR.html), [VkVideoEncodeH265SessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionCreateInfoKHR.html), [VkVideoEncodeSessionIntraRefreshCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeSessionIntraRefreshCreateInfoKHR.html), or [VkVideoEncodeSessionRgbConversionCreateInfoVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeSessionRgbConversionCreateInfoVALVE.html)
- [](#VUID-VkVideoSessionCreateInfoKHR-sType-unique)VUID-VkVideoSessionCreateInfoKHR-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkVideoSessionCreateInfoKHR-flags-parameter)VUID-VkVideoSessionCreateInfoKHR-flags-parameter  
  `flags` **must** be a valid combination of [VkVideoSessionCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionCreateFlagBitsKHR.html) values
- [](#VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-parameter)VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-parameter  
  `pVideoProfile` **must** be a valid pointer to a valid [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html) structure
- [](#VUID-VkVideoSessionCreateInfoKHR-pictureFormat-parameter)VUID-VkVideoSessionCreateInfoKHR-pictureFormat-parameter  
  `pictureFormat` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value
- [](#VUID-VkVideoSessionCreateInfoKHR-referencePictureFormat-parameter)VUID-VkVideoSessionCreateInfoKHR-referencePictureFormat-parameter  
  `referencePictureFormat` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value
- [](#VUID-VkVideoSessionCreateInfoKHR-pStdHeaderVersion-parameter)VUID-VkVideoSessionCreateInfoKHR-pStdHeaderVersion-parameter  
  `pStdHeaderVersion` **must** be a valid pointer to a valid [VkExtensionProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtensionProperties.html) structure

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkExtensionProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtensionProperties.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html), [VkVideoSessionCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionCreateFlagsKHR.html), [vkCreateVideoSessionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateVideoSessionKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoSessionCreateInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0