# vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT(3) Manual Page

## Name

vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT - Get
acceleration structure opaque capture descriptor data



## <a href="#_c_specification" class="anchor"></a>C Specification

To get the opaque capture descriptor data for an acceleration structure,
call:

``` c
// Provided by VK_EXT_descriptor_buffer with VK_KHR_acceleration_structure or VK_NV_ray_tracing
VkResult vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT(
    VkDevice                                    device,
    const VkAccelerationStructureCaptureDescriptorDataInfoEXT* pInfo,
    void*                                       pData);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that gets the data.

- `pInfo` is a pointer to a
  [VkAccelerationStructureCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCaptureDescriptorDataInfoEXT.html)
  structure specifying the acceleration structure.

- `pData` is a pointer to an application-allocated buffer where the data
  will be written.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-None-08088"
  id="VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-None-08088"></a>
  VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-None-08088  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-descriptorBuffer"
  target="_blank"
  rel="noopener"><code>descriptorBufferCaptureReplay</code></a> feature
  **must** be enabled

- <a
  href="#VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-pData-08089"
  id="VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-pData-08089"></a>
  VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-pData-08089  
  `pData` **must** point to a buffer that is at least
  [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`accelerationStructureCaptureReplayDescriptorDataSize`
  bytes in size

- <a
  href="#VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-device-08090"
  id="VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-device-08090"></a>
  VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-device-08090  
  If `device` was created with multiple physical devices, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-bufferDeviceAddressMultiDevice"
  target="_blank"
  rel="noopener"><code>bufferDeviceAddressMultiDevice</code></a> feature
  **must** be enabled

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-device-parameter"
  id="VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-device-parameter"></a>
  VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-pInfo-parameter"
  id="VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-pInfo-parameter"></a>
  VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkAccelerationStructureCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCaptureDescriptorDataInfoEXT.html)
  structure

- <a
  href="#VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-pData-parameter"
  id="VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-pData-parameter"></a>
  VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-pData-parameter  
  `pData` **must** be a pointer value

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_descriptor_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_buffer.html),
[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkAccelerationStructureCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCaptureDescriptorDataInfoEXT.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
