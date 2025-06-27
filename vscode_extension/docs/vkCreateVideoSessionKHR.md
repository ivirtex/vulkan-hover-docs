# vkCreateVideoSessionKHR(3) Manual Page

## Name

vkCreateVideoSessionKHR - Creates a video session object



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a video session object, call:

``` c
// Provided by VK_KHR_video_queue
VkResult vkCreateVideoSessionKHR(
    VkDevice                                    device,
    const VkVideoSessionCreateInfoKHR*          pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkVideoSessionKHR*                          pVideoSession);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the video session.

- `pCreateInfo` is a pointer to a
  [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)
  structure containing parameters to be used to create the video
  session.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pVideoSession` is a pointer to a
  [VkVideoSessionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionKHR.html) handle in which the
  resulting video session object is returned.

## <a href="#_description" class="anchor"></a>Description

The resulting video session object is said to be created with the video
codec operation specified in
`pCreateInfo->pVideoProfile->videoCodecOperation`.

The name and version of the codec-specific Video Std header to be used
with the video session is specified by the
[VkExtensionProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtensionProperties.html) structure pointed to
by `pCreateInfo->pStdHeaderVersion`. If a non-existent or unsupported
Video Std header version is specified in
`pCreateInfo->pStdHeaderVersion->specVersion`, then this command returns
`VK_ERROR_VIDEO_STD_VERSION_NOT_SUPPORTED_KHR`.

Video session objects are created in *uninitialized* state. In order to
transition the video session into *initial* state, the application
**must** issue a
[vkCmdControlVideoCodingKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdControlVideoCodingKHR.html) command
with
[VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodingControlInfoKHR.html)::`flags`
including `VK_VIDEO_CODING_CONTROL_RESET_BIT_KHR`.

Video session objects also maintain the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-state-and-backing-store"
target="_blank" rel="noopener">state</a> of the DPB. The number of DPB
slots usable with the created video session is specified in
`pCreateInfo->maxDpbSlots`, and each slot is initially in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot-states"
target="_blank" rel="noopener">inactive state</a>.

Each <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
target="_blank" rel="noopener">DPB slot</a> maintained by the created
video session **can** refer to a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#reference-picture"
target="_blank" rel="noopener">reference picture</a> representing a
video frame.

In addition, if the `videoCodecOperation` member of the
[VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html) structure pointed to
by `pCreateInfo->pVideoProfile` is
`VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR` and the `pictureLayout`
member of the
[VkVideoDecodeH264ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264ProfileInfoKHR.html)
structure provided in the
[VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html)::`pNext` chain is
not `VK_VIDEO_DECODE_H264_PICTURE_LAYOUT_PROGRESSIVE_KHR`, then the
created video session supports *interlaced* frames and each <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
target="_blank" rel="noopener">DPB slot</a> maintained by the created
video session **can** instead refer to separate top field and bottom
field <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#reference-picture"
target="_blank" rel="noopener">reference pictures</a> that together
**can** represent a full video frame. In this case, it is up to the
application, driven by the video content, whether it associates any
individual DPB slot with separate top and/or bottom field pictures or a
single picture representing a full frame.

The created video session **can** be used to perform video coding
operations using video frames up to the maximum size specified in
`pCreateInfo->maxCodedExtent`. The minimum frame size allowed is
implicitly derived from
[VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html)::`minCodedExtent`,
as returned by
[vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
for the video profile specified by `pCreateInfo->pVideoProfile`.
Accordingly, the created video session is said to be created with a
`minCodedExtent` equal to that.

In case of video session objects created with a video encode operation,
implementations **may** return the
`VK_ERROR_INVALID_VIDEO_STD_PARAMETERS_KHR` error if any of the
specified Video Std parameters do not adhere to the syntactic or
semantic requirements of the used video compression standard, or if
values derived from parameters according to the rules defined by the
used video compression standard do not adhere to the capabilities of the
video compression standard or the implementation.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Applications <strong>should</strong> not rely on the
<code>VK_ERROR_INVALID_VIDEO_STD_PARAMETERS_KHR</code> error being
returned by any command as a means to verify Video Std parameters, as
implementations are not required to report the error in any specific set
of cases.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-vkCreateVideoSessionKHR-device-parameter"
  id="VUID-vkCreateVideoSessionKHR-device-parameter"></a>
  VUID-vkCreateVideoSessionKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateVideoSessionKHR-pCreateInfo-parameter"
  id="VUID-vkCreateVideoSessionKHR-pCreateInfo-parameter"></a>
  VUID-vkCreateVideoSessionKHR-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)
  structure

- <a href="#VUID-vkCreateVideoSessionKHR-pAllocator-parameter"
  id="VUID-vkCreateVideoSessionKHR-pAllocator-parameter"></a>
  VUID-vkCreateVideoSessionKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateVideoSessionKHR-pVideoSession-parameter"
  id="VUID-vkCreateVideoSessionKHR-pVideoSession-parameter"></a>
  VUID-vkCreateVideoSessionKHR-pVideoSession-parameter  
  `pVideoSession` **must** be a valid pointer to a
  [VkVideoSessionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionKHR.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_INITIALIZATION_FAILED`

- `VK_ERROR_VIDEO_STD_VERSION_NOT_SUPPORTED_KHR`

- `VK_ERROR_INVALID_VIDEO_STD_PARAMETERS_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html),
[VkVideoSessionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateVideoSessionKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
