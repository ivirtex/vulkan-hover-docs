# vkGetMemoryAndroidHardwareBufferANDROID(3) Manual Page

## Name

vkGetMemoryAndroidHardwareBufferANDROID - Get an Android hardware buffer for a memory object



## [](#_c_specification)C Specification

To export an Android hardware buffer referencing the payload of a Vulkan device memory object, call:

```c++
// Provided by VK_ANDROID_external_memory_android_hardware_buffer
VkResult vkGetMemoryAndroidHardwareBufferANDROID(
    VkDevice                                    device,
    const VkMemoryGetAndroidHardwareBufferInfoANDROID* pInfo,
    struct AHardwareBuffer**                    pBuffer);
```

## [](#_parameters)Parameters

- `device` is the logical device that created the device memory being exported.
- `pInfo` is a pointer to a [VkMemoryGetAndroidHardwareBufferInfoANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetAndroidHardwareBufferInfoANDROID.html) structure containing parameters of the export operation.
- `pBuffer` will return an Android hardware buffer referencing the payload of the device memory object.

## [](#_description)Description

Each call to `vkGetMemoryAndroidHardwareBufferANDROID` **must** return an Android hardware buffer with a new reference acquired in addition to the reference held by the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html). To avoid leaking resources, the application **must** release the reference by calling `AHardwareBuffer_release` when it is no longer needed. When called with the same handle in [VkMemoryGetAndroidHardwareBufferInfoANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetAndroidHardwareBufferInfoANDROID.html)::`memory`, `vkGetMemoryAndroidHardwareBufferANDROID` **must** return the same Android hardware buffer object. If the device memory was created by importing an Android hardware buffer, `vkGetMemoryAndroidHardwareBufferANDROID` **must** return that same Android hardware buffer object.

Valid Usage (Implicit)

- [](#VUID-vkGetMemoryAndroidHardwareBufferANDROID-device-parameter)VUID-vkGetMemoryAndroidHardwareBufferANDROID-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetMemoryAndroidHardwareBufferANDROID-pInfo-parameter)VUID-vkGetMemoryAndroidHardwareBufferANDROID-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkMemoryGetAndroidHardwareBufferInfoANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetAndroidHardwareBufferInfoANDROID.html) structure
- [](#VUID-vkGetMemoryAndroidHardwareBufferANDROID-pBuffer-parameter)VUID-vkGetMemoryAndroidHardwareBufferANDROID-pBuffer-parameter  
  `pBuffer` **must** be a valid pointer to a valid pointer to an [AHardwareBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/AHardwareBuffer.html) value

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_TOO_MANY_OBJECTS`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_ANDROID\_external\_memory\_android\_hardware\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ANDROID_external_memory_android_hardware_buffer.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkMemoryGetAndroidHardwareBufferInfoANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetAndroidHardwareBufferInfoANDROID.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetMemoryAndroidHardwareBufferANDROID)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0