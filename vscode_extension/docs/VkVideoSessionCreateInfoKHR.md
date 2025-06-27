# VkVideoSessionCreateInfoKHR(3) Manual Page

## Name

VkVideoSessionCreateInfoKHR - Structure specifying parameters of a newly
created video session



## <a href="#_c_specification" class="anchor"></a>C Specification

The [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)
structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `queueFamilyIndex` is the index of the queue family the created video
  session will be used with.

- `flags` is a bitmask of
  [VkVideoSessionCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateFlagBitsKHR.html)
  specifying creation flags.

- `pVideoProfile` is a pointer to a
  [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html) structure
  specifying the video profile the created video session will be used
  with.

- `pictureFormat` is the image format the created video session will be
  used with. If `pVideoProfile->videoCodecOperation` specifies a decode
  operation, then `pictureFormat` is the image format of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-output-picture"
  target="_blank" rel="noopener">decode output pictures</a> usable with
  the created video session. If `pVideoProfile->videoCodecOperation`
  specifies an encode operation, then `pictureFormat` is the image
  format of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-input-picture"
  target="_blank" rel="noopener">encode input pictures</a> usable with
  the created video session.

- `maxCodedExtent` is the maximum width and height of the coded frames
  the created video session will be used with.

- `referencePictureFormat` is the image format of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#reference-picture"
  target="_blank" rel="noopener">reference pictures</a> stored in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb"
  target="_blank" rel="noopener">DPB</a> the created video session will
  be used with.

- `maxDpbSlots` is the maximum number of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
  target="_blank" rel="noopener">DPB Slots</a> that **can** be used with
  the created video session.

- `maxActiveReferencePictures` is the maximum number of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#active-reference-pictures"
  target="_blank" rel="noopener">active reference pictures</a> that
  **can** be used in a single video coding operation using the created
  video session.

- `pStdHeaderVersion` is a pointer to a
  [VkExtensionProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtensionProperties.html) structure
  requesting the Video Std header version to use for the
  `videoCodecOperation` specified in `pVideoProfile`.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkVideoSessionCreateInfoKHR-protectedMemory-07189"
  id="VUID-VkVideoSessionCreateInfoKHR-protectedMemory-07189"></a>
  VUID-VkVideoSessionCreateInfoKHR-protectedMemory-07189  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-protectedMemory"
  target="_blank" rel="noopener"><code>protectedMemory</code></a>
  feature is not enabled or if
  [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html)::`flags` does
  not include `VK_VIDEO_CAPABILITY_PROTECTED_CONTENT_BIT_KHR`, as
  returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile specified by `pVideoProfile`, then `flags`
  **must** not include
  `VK_VIDEO_SESSION_CREATE_PROTECTED_CONTENT_BIT_KHR`

- <a href="#VUID-VkVideoSessionCreateInfoKHR-flags-08371"
  id="VUID-VkVideoSessionCreateInfoKHR-flags-08371"></a>
  VUID-VkVideoSessionCreateInfoKHR-flags-08371  
  If `flags` includes `VK_VIDEO_SESSION_CREATE_INLINE_QUERIES_BIT_KHR`,
  then <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-videoMaintenance1"
  target="_blank" rel="noopener"><code>videoMaintenance1</code></a>
  **must** be enabled

- <a href="#VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-04845"
  id="VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-04845"></a>
  VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-04845  
  `pVideoProfile` **must** be a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-profile-support"
  target="_blank" rel="noopener">supported video profile</a>

- <a href="#VUID-VkVideoSessionCreateInfoKHR-maxDpbSlots-04847"
  id="VUID-VkVideoSessionCreateInfoKHR-maxDpbSlots-04847"></a>
  VUID-VkVideoSessionCreateInfoKHR-maxDpbSlots-04847  
  `maxDpbSlots` **must** be less than or equal to
  [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html)::`maxDpbSlots`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile specified by `pVideoProfile`

- <a
  href="#VUID-VkVideoSessionCreateInfoKHR-maxActiveReferencePictures-04849"
  id="VUID-VkVideoSessionCreateInfoKHR-maxActiveReferencePictures-04849"></a>
  VUID-VkVideoSessionCreateInfoKHR-maxActiveReferencePictures-04849  
  `maxActiveReferencePictures` **must** be less than or equal to
  [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html)::`maxActiveReferencePictures`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile specified by `pVideoProfile`

- <a href="#VUID-VkVideoSessionCreateInfoKHR-maxDpbSlots-04850"
  id="VUID-VkVideoSessionCreateInfoKHR-maxDpbSlots-04850"></a>
  VUID-VkVideoSessionCreateInfoKHR-maxDpbSlots-04850  
  If either `maxDpbSlots` or `maxActiveReferencePictures` is `0`, then
  both **must** be `0`

- <a href="#VUID-VkVideoSessionCreateInfoKHR-maxCodedExtent-04851"
  id="VUID-VkVideoSessionCreateInfoKHR-maxCodedExtent-04851"></a>
  VUID-VkVideoSessionCreateInfoKHR-maxCodedExtent-04851  
  `maxCodedExtent` **must** be between
  [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html)::`minCodedExtent`
  and
  [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html)::`maxCodedExtent`,
  inclusive, as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile specified by `pVideoProfile`

- <a href="#VUID-VkVideoSessionCreateInfoKHR-referencePictureFormat-04852"
  id="VUID-VkVideoSessionCreateInfoKHR-referencePictureFormat-04852"></a>
  VUID-VkVideoSessionCreateInfoKHR-referencePictureFormat-04852  
  If `pVideoProfile->videoCodecOperation` specifies a decode operation
  and `maxActiveReferencePictures` is greater than `0`, then
  `referencePictureFormat` **must** be one of the supported decode DPB
  formats, as returned by
  [vkGetPhysicalDeviceVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoFormatPropertiesKHR.html)
  in
  [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoFormatPropertiesKHR.html)::`format`
  when called with the `imageUsage` member of its `pVideoFormatInfo`
  parameter containing `VK_IMAGE_USAGE_VIDEO_DECODE_DPB_BIT_KHR`, and
  with a [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html)
  structure specified in the `pNext` chain of its `pVideoFormatInfo`
  parameter whose `pProfiles` member contains an element matching
  `pVideoProfile`

- <a href="#VUID-VkVideoSessionCreateInfoKHR-referencePictureFormat-06814"
  id="VUID-VkVideoSessionCreateInfoKHR-referencePictureFormat-06814"></a>
  VUID-VkVideoSessionCreateInfoKHR-referencePictureFormat-06814  
  If `pVideoProfile->videoCodecOperation` specifies an encode operation
  and `maxActiveReferencePictures` is greater than `0`, then
  `referencePictureFormat` **must** be one of the supported decode DPB
  formats, as returned by then `referencePictureFormat` **must** be one
  of the supported encode DPB formats, as returned by
  [vkGetPhysicalDeviceVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoFormatPropertiesKHR.html)
  in
  [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoFormatPropertiesKHR.html)::`format`
  when called with the `imageUsage` member of its `pVideoFormatInfo`
  parameter containing `VK_IMAGE_USAGE_VIDEO_ENCODE_DPB_BIT_KHR`, and
  with a [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html)
  structure specified in the `pNext` chain of its `pVideoFormatInfo`
  parameter whose `pProfiles` member contains an element matching
  `pVideoProfile`

- <a href="#VUID-VkVideoSessionCreateInfoKHR-pictureFormat-04853"
  id="VUID-VkVideoSessionCreateInfoKHR-pictureFormat-04853"></a>
  VUID-VkVideoSessionCreateInfoKHR-pictureFormat-04853  
  If `pVideoProfile->videoCodecOperation` specifies a decode operation,
  then `pictureFormat` **must** be one of the supported decode output
  formats, as returned by
  [vkGetPhysicalDeviceVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoFormatPropertiesKHR.html)
  in
  [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoFormatPropertiesKHR.html)::`format`
  when called with the `imageUsage` member of its `pVideoFormatInfo`
  parameter containing `VK_IMAGE_USAGE_VIDEO_DECODE_DST_BIT_KHR`, and
  with a [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html)
  structure specified in the `pNext` chain of its `pVideoFormatInfo`
  parameter whose `pProfiles` member contains an element matching
  `pVideoProfile`

- <a href="#VUID-VkVideoSessionCreateInfoKHR-pictureFormat-04854"
  id="VUID-VkVideoSessionCreateInfoKHR-pictureFormat-04854"></a>
  VUID-VkVideoSessionCreateInfoKHR-pictureFormat-04854  
  If `pVideoProfile->videoCodecOperation` specifies an encode operation,
  then `pictureFormat` **must** be one of the supported encode input
  formats, as returned by
  [vkGetPhysicalDeviceVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoFormatPropertiesKHR.html)
  in
  [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoFormatPropertiesKHR.html)::`format`
  when called with the `imageUsage` member of its `pVideoFormatInfo`
  parameter containing `VK_IMAGE_USAGE_VIDEO_ENCODE_SRC_BIT_KHR`, and
  with a [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html)
  structure specified in the `pNext` chain of its `pVideoFormatInfo`
  parameter whose `pProfiles` member contains an element matching
  `pVideoProfile`

- <a href="#VUID-VkVideoSessionCreateInfoKHR-pStdHeaderVersion-07190"
  id="VUID-VkVideoSessionCreateInfoKHR-pStdHeaderVersion-07190"></a>
  VUID-VkVideoSessionCreateInfoKHR-pStdHeaderVersion-07190  
  `pStdHeaderVersion->extensionName` **must** match
  [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html)::`stdHeaderVersion.extensionName`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile specified by `pVideoProfile`

- <a href="#VUID-VkVideoSessionCreateInfoKHR-pStdHeaderVersion-07191"
  id="VUID-VkVideoSessionCreateInfoKHR-pStdHeaderVersion-07191"></a>
  VUID-VkVideoSessionCreateInfoKHR-pStdHeaderVersion-07191  
  `pStdHeaderVersion->specVersion` **must** be less than or equal to
  [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html)::`stdHeaderVersion.specVersion`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile specified by `pVideoProfile`

- <a href="#VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-08251"
  id="VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-08251"></a>
  VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-08251  
  If `pVideoProfile->videoCodecOperation` is
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR` and the `pNext` chain
  of this structure includes a
  [VkVideoEncodeH264SessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionCreateInfoKHR.html)
  structure, then its `maxLevelIdc` member **must** be less than or
  equal to
  [VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilitiesKHR.html)::`maxLevelIdc`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile specified in `pVideoProfile`

- <a href="#VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-08252"
  id="VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-08252"></a>
  VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-08252  
  If `pVideoProfile->videoCodecOperation` is
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR` and the `pNext` chain
  of this structure includes a
  [VkVideoEncodeH265SessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionCreateInfoKHR.html)
  structure, then its `maxLevelIdc` member **must** be less than or
  equal to
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`maxLevelIdc`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile specified in `pVideoProfile`

Valid Usage (Implicit)

- <a href="#VUID-VkVideoSessionCreateInfoKHR-sType-sType"
  id="VUID-VkVideoSessionCreateInfoKHR-sType-sType"></a>
  VUID-VkVideoSessionCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_SESSION_CREATE_INFO_KHR`

- <a href="#VUID-VkVideoSessionCreateInfoKHR-pNext-pNext"
  id="VUID-VkVideoSessionCreateInfoKHR-pNext-pNext"></a>
  VUID-VkVideoSessionCreateInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkVideoEncodeH264SessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionCreateInfoKHR.html)
  or
  [VkVideoEncodeH265SessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionCreateInfoKHR.html)

- <a href="#VUID-VkVideoSessionCreateInfoKHR-sType-unique"
  id="VUID-VkVideoSessionCreateInfoKHR-sType-unique"></a>
  VUID-VkVideoSessionCreateInfoKHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkVideoSessionCreateInfoKHR-flags-parameter"
  id="VUID-VkVideoSessionCreateInfoKHR-flags-parameter"></a>
  VUID-VkVideoSessionCreateInfoKHR-flags-parameter  
  `flags` **must** be a valid combination of
  [VkVideoSessionCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateFlagBitsKHR.html)
  values

- <a href="#VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-parameter"
  id="VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-parameter"></a>
  VUID-VkVideoSessionCreateInfoKHR-pVideoProfile-parameter  
  `pVideoProfile` **must** be a valid pointer to a valid
  [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html) structure

- <a href="#VUID-VkVideoSessionCreateInfoKHR-pictureFormat-parameter"
  id="VUID-VkVideoSessionCreateInfoKHR-pictureFormat-parameter"></a>
  VUID-VkVideoSessionCreateInfoKHR-pictureFormat-parameter  
  `pictureFormat` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

- <a
  href="#VUID-VkVideoSessionCreateInfoKHR-referencePictureFormat-parameter"
  id="VUID-VkVideoSessionCreateInfoKHR-referencePictureFormat-parameter"></a>
  VUID-VkVideoSessionCreateInfoKHR-referencePictureFormat-parameter  
  `referencePictureFormat` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html)
  value

- <a href="#VUID-VkVideoSessionCreateInfoKHR-pStdHeaderVersion-parameter"
  id="VUID-VkVideoSessionCreateInfoKHR-pStdHeaderVersion-parameter"></a>
  VUID-VkVideoSessionCreateInfoKHR-pStdHeaderVersion-parameter  
  `pStdHeaderVersion` **must** be a valid pointer to a valid
  [VkExtensionProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtensionProperties.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkExtensionProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtensionProperties.html),
[VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html), [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html),
[VkVideoSessionCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateFlagsKHR.html),
[vkCreateVideoSessionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateVideoSessionKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoSessionCreateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
