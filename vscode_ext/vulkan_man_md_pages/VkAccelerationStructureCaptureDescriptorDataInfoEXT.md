# VkAccelerationStructureCaptureDescriptorDataInfoEXT(3) Manual Page

## Name

VkAccelerationStructureCaptureDescriptorDataInfoEXT - Structure
specifying an acceleration structure for descriptor capture



## <a href="#_c_specification" class="anchor"></a>C Specification

Information about the acceleration structure to get descriptor buffer
capture data for is passed in a
`VkAccelerationStructureCaptureDescriptorDataInfoEXT` structure:

``` c
// Provided by VK_EXT_descriptor_buffer with VK_KHR_acceleration_structure or VK_NV_ray_tracing
typedef struct VkAccelerationStructureCaptureDescriptorDataInfoEXT {
    VkStructureType               sType;
    const void*                   pNext;
    VkAccelerationStructureKHR    accelerationStructure;
    VkAccelerationStructureNV     accelerationStructureNV;
} VkAccelerationStructureCaptureDescriptorDataInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `accelerationStructure` is the `VkAccelerationStructureKHR` handle of
  the acceleration structure to get opaque capture data for.

- `accelerationStructureNV` is the `VkAccelerationStructureNV` handle of
  the acceleration structure to get opaque capture data for.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructure-08091"
  id="VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructure-08091"></a>
  VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructure-08091  
  If `accelerationStructure` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) then `accelerationStructure`
  **must** have been created with
  `VK_ACCELERATION_STRUCTURE_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`
  set in
  [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoKHR.html)::`createFlags`

- <a
  href="#VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructureNV-08092"
  id="VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructureNV-08092"></a>
  VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructureNV-08092  
  If `accelerationStructureNV` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) then `accelerationStructureNV`
  **must** have been created with
  `VK_ACCELERATION_STRUCTURE_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`
  set in
  [VkAccelerationStructureCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoNV.html)::`info.flags`

- <a
  href="#VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructure-08093"
  id="VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructure-08093"></a>
  VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructure-08093  
  If `accelerationStructure` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) then `accelerationStructureNV`
  **must** be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a
  href="#VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructureNV-08094"
  id="VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructureNV-08094"></a>
  VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructureNV-08094  
  If `accelerationStructureNV` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) then `accelerationStructure`
  **must** be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

Valid Usage (Implicit)

- <a
  href="#VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-sType-sType"
  id="VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-sType-sType"></a>
  VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_CAPTURE_DESCRIPTOR_DATA_INFO_EXT`

- <a
  href="#VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-pNext-pNext"
  id="VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-pNext-pNext"></a>
  VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a
  href="#VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructure-parameter"
  id="VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructure-parameter"></a>
  VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructure-parameter  
  If `accelerationStructure` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `accelerationStructure`
  **must** be a valid
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html) handle

- <a
  href="#VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructureNV-parameter"
  id="VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructureNV-parameter"></a>
  VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructureNV-parameter  
  If `accelerationStructureNV` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `accelerationStructureNV`
  **must** be a valid
  [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html) handle

- <a
  href="#VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-commonparent"
  id="VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-commonparent"></a>
  VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-commonparent  
  Both of `accelerationStructure`, and `accelerationStructureNV` that
  are valid handles of non-ignored parameters **must** have been
  created, allocated, or retrieved from the same
  [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_descriptor_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_buffer.html),
[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html),
[VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAccelerationStructureCaptureDescriptorDataInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
