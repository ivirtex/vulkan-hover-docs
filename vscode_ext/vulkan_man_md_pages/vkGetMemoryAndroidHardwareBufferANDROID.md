# vkGetMemoryAndroidHardwareBufferANDROID(3) Manual Page

## Name

vkGetMemoryAndroidHardwareBufferANDROID - Get an Android hardware buffer
for a memory object



## <a href="#_c_specification" class="anchor"></a>C Specification

To export an Android hardware buffer referencing the payload of a Vulkan
device memory object, call:

``` c
// Provided by VK_ANDROID_external_memory_android_hardware_buffer
VkResult vkGetMemoryAndroidHardwareBufferANDROID(
    VkDevice                                    device,
    const VkMemoryGetAndroidHardwareBufferInfoANDROID* pInfo,
    struct AHardwareBuffer**                    pBuffer);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that created the device memory being
  exported.

- `pInfo` is a pointer to a
  [VkMemoryGetAndroidHardwareBufferInfoANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryGetAndroidHardwareBufferInfoANDROID.html)
  structure containing parameters of the export operation.

- `pBuffer` will return an Android hardware buffer referencing the
  payload of the device memory object.

## <a href="#_description" class="anchor"></a>Description

Each call to `vkGetMemoryAndroidHardwareBufferANDROID` **must** return
an Android hardware buffer with a new reference acquired in addition to
the reference held by the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html). To
avoid leaking resources, the application **must** release the reference
by calling `AHardwareBuffer_release` when it is no longer needed. When
called with the same handle in
[VkMemoryGetAndroidHardwareBufferInfoANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryGetAndroidHardwareBufferInfoANDROID.html)::`memory`,
`vkGetMemoryAndroidHardwareBufferANDROID` **must** return the same
Android hardware buffer object. If the device memory was created by
importing an Android hardware buffer,
`vkGetMemoryAndroidHardwareBufferANDROID` **must** return that same
Android hardware buffer object.

Valid Usage (Implicit)

- <a href="#VUID-vkGetMemoryAndroidHardwareBufferANDROID-device-parameter"
  id="VUID-vkGetMemoryAndroidHardwareBufferANDROID-device-parameter"></a>
  VUID-vkGetMemoryAndroidHardwareBufferANDROID-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetMemoryAndroidHardwareBufferANDROID-pInfo-parameter"
  id="VUID-vkGetMemoryAndroidHardwareBufferANDROID-pInfo-parameter"></a>
  VUID-vkGetMemoryAndroidHardwareBufferANDROID-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkMemoryGetAndroidHardwareBufferInfoANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryGetAndroidHardwareBufferInfoANDROID.html)
  structure

- <a
  href="#VUID-vkGetMemoryAndroidHardwareBufferANDROID-pBuffer-parameter"
  id="VUID-vkGetMemoryAndroidHardwareBufferANDROID-pBuffer-parameter"></a>
  VUID-vkGetMemoryAndroidHardwareBufferANDROID-pBuffer-parameter  
  `pBuffer` **must** be a valid pointer to a valid pointer to an
  [AHardwareBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/AHardwareBuffer.html) value

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_TOO_MANY_OBJECTS`

- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_ANDROID_external_memory_android_hardware_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_ANDROID_external_memory_android_hardware_buffer.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkMemoryGetAndroidHardwareBufferInfoANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryGetAndroidHardwareBufferInfoANDROID.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetMemoryAndroidHardwareBufferANDROID"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
