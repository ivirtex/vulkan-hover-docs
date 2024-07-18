# vkGetAccelerationStructureDeviceAddressKHR(3) Manual Page

## Name

vkGetAccelerationStructureDeviceAddressKHR - Query an address of an
acceleration structure



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the 64-bit device address for an acceleration structure, call:

``` c
// Provided by VK_KHR_acceleration_structure
VkDeviceAddress vkGetAccelerationStructureDeviceAddressKHR(
    VkDevice                                    device,
    const VkAccelerationStructureDeviceAddressInfoKHR* pInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that the acceleration structure was
  created on.

- `pInfo` is a pointer to a
  [VkAccelerationStructureDeviceAddressInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureDeviceAddressInfoKHR.html)
  structure specifying the acceleration structure to retrieve an address
  for.

## <a href="#_description" class="anchor"></a>Description

The 64-bit return value is an address of the acceleration structure,
which can be used for device and shader operations that involve
acceleration structures, such as ray traversal and acceleration
structure building.

If the acceleration structure was created with a non-zero value of
[VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoKHR.html)::`deviceAddress`,
the return value will be the same address.

If the acceleration structure was created with a `type` of
`VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR`, the returned address
**must** be consistent with the relative offset to other acceleration
structures with `type` `VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR`
allocated with the same [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html). That is, the
difference in returned addresses between the two **must** be the same as
the difference in offsets provided at acceleration structure creation.

The returned address **must** be aligned to 256 bytes.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>The acceleration structure device address <strong>may</strong> be
different from the buffer device address corresponding to the
acceleration structure’s start offset in its storage buffer for
acceleration structure types other than
<code>VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR</code>.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a
  href="#VUID-vkGetAccelerationStructureDeviceAddressKHR-accelerationStructure-08935"
  id="VUID-vkGetAccelerationStructureDeviceAddressKHR-accelerationStructure-08935"></a>
  VUID-vkGetAccelerationStructureDeviceAddressKHR-accelerationStructure-08935  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-accelerationStructure"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceAccelerationStructureFeaturesKHR</code>::<code>accelerationStructure</code></a>
  feature **must** be enabled

- <a href="#VUID-vkGetAccelerationStructureDeviceAddressKHR-device-03504"
  id="VUID-vkGetAccelerationStructureDeviceAddressKHR-device-03504"></a>
  VUID-vkGetAccelerationStructureDeviceAddressKHR-device-03504  
  If `device` was created with multiple physical devices, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-bufferDeviceAddressMultiDevice"
  target="_blank"
  rel="noopener"><code>bufferDeviceAddressMultiDevice</code></a> feature
  **must** be enabled

- <a href="#VUID-vkGetAccelerationStructureDeviceAddressKHR-pInfo-09541"
  id="VUID-vkGetAccelerationStructureDeviceAddressKHR-pInfo-09541"></a>
  VUID-vkGetAccelerationStructureDeviceAddressKHR-pInfo-09541  
  If the buffer on which `pInfo->accelerationStructure` was placed is
  non-sparse then it **must** be bound completely and contiguously to a
  single `VkDeviceMemory` object

- <a href="#VUID-vkGetAccelerationStructureDeviceAddressKHR-pInfo-09542"
  id="VUID-vkGetAccelerationStructureDeviceAddressKHR-pInfo-09542"></a>
  VUID-vkGetAccelerationStructureDeviceAddressKHR-pInfo-09542  
  The buffer on which `pInfo->accelerationStructure` was placed **must**
  have been created with the `VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT`
  usage flag

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetAccelerationStructureDeviceAddressKHR-device-parameter"
  id="VUID-vkGetAccelerationStructureDeviceAddressKHR-device-parameter"></a>
  VUID-vkGetAccelerationStructureDeviceAddressKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetAccelerationStructureDeviceAddressKHR-pInfo-parameter"
  id="VUID-vkGetAccelerationStructureDeviceAddressKHR-pInfo-parameter"></a>
  VUID-vkGetAccelerationStructureDeviceAddressKHR-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkAccelerationStructureDeviceAddressInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureDeviceAddressInfoKHR.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkAccelerationStructureDeviceAddressInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureDeviceAddressInfoKHR.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetAccelerationStructureDeviceAddressKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
