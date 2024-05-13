# vkCreateAccelerationStructureKHR(3) Manual Page

## Name

vkCreateAccelerationStructureKHR - Create a new acceleration structure
object



## <a href="#_c_specification" class="anchor"></a>C Specification

To create an acceleration structure, call:

``` c
// Provided by VK_KHR_acceleration_structure
VkResult vkCreateAccelerationStructureKHR(
    VkDevice                                    device,
    const VkAccelerationStructureCreateInfoKHR* pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkAccelerationStructureKHR*                 pAccelerationStructure);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the acceleration structure
  object.

- `pCreateInfo` is a pointer to a
  [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoKHR.html)
  structure containing parameters affecting creation of the acceleration
  structure.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pAccelerationStructure` is a pointer to a
  `VkAccelerationStructureKHR` handle in which the resulting
  acceleration structure object is returned.

## <a href="#_description" class="anchor"></a>Description

Similar to other objects in Vulkan, the acceleration structure creation
merely creates an object with a specific “shape”. The type and quantity
of geometry that can be built into an acceleration structure is
determined by the parameters of
[VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoKHR.html).

The acceleration structure data is stored in the object referred to by
`VkAccelerationStructureCreateInfoKHR`::`buffer`. Once memory has been
bound to that buffer, it **must** be populated by acceleration structure
build or acceleration structure copy commands such as
[vkCmdBuildAccelerationStructuresKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBuildAccelerationStructuresKHR.html),
[vkBuildAccelerationStructuresKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBuildAccelerationStructuresKHR.html),
[vkCmdCopyAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyAccelerationStructureKHR.html),
and
[vkCopyAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyAccelerationStructureKHR.html).

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The expected usage for a trace capture/replay tool is that it will
serialize and later deserialize the acceleration structure data using
acceleration structure copy commands. During capture the tool will use
<a
href="vkCopyAccelerationStructureToMemoryKHR.html">vkCopyAccelerationStructureToMemoryKHR</a>
or <a
href="vkCmdCopyAccelerationStructureToMemoryKHR.html">vkCmdCopyAccelerationStructureToMemoryKHR</a>
with a <code>mode</code> of
<code>VK_COPY_ACCELERATION_STRUCTURE_MODE_SERIALIZE_KHR</code>, and <a
href="vkCopyMemoryToAccelerationStructureKHR.html">vkCopyMemoryToAccelerationStructureKHR</a>
or <a
href="vkCmdCopyMemoryToAccelerationStructureKHR.html">vkCmdCopyMemoryToAccelerationStructureKHR</a>
with a <code>mode</code> of
<code>VK_COPY_ACCELERATION_STRUCTURE_MODE_DESERIALIZE_KHR</code> during
replay.</p></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Memory does not need to be bound to the underlying buffer when <a
href="vkCreateAccelerationStructureKHR.html">vkCreateAccelerationStructureKHR</a>
is called.</p></td>
</tr>
</tbody>
</table>

The input buffers passed to acceleration structure build commands will
be referenced by the implementation for the duration of the command.
After the command completes, the acceleration structure **may** hold a
reference to any acceleration structure specified by an active instance
contained therein. Apart from this referencing, acceleration structures
**must** be fully self-contained. The application **can** reuse or free
any memory which was used by the command as an input or as scratch
without affecting the results of ray traversal.

Valid Usage

- <a
  href="#VUID-vkCreateAccelerationStructureKHR-accelerationStructure-03611"
  id="VUID-vkCreateAccelerationStructureKHR-accelerationStructure-03611"></a>
  VUID-vkCreateAccelerationStructureKHR-accelerationStructure-03611  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-accelerationStructure"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceAccelerationStructureFeaturesKHR</code>::<code>accelerationStructure</code></a>
  feature **must** be enabled

- <a href="#VUID-vkCreateAccelerationStructureKHR-deviceAddress-03488"
  id="VUID-vkCreateAccelerationStructureKHR-deviceAddress-03488"></a>
  VUID-vkCreateAccelerationStructureKHR-deviceAddress-03488  
  If
  [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoKHR.html)::`deviceAddress`
  is not zero, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-accelerationStructureCaptureReplay"
  target="_blank"
  rel="noopener"><code>accelerationStructureCaptureReplay</code></a>
  feature **must** be enabled

- <a href="#VUID-vkCreateAccelerationStructureKHR-device-03489"
  id="VUID-vkCreateAccelerationStructureKHR-device-03489"></a>
  VUID-vkCreateAccelerationStructureKHR-device-03489  
  If `device` was created with multiple physical devices, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-bufferDeviceAddressMultiDevice"
  target="_blank"
  rel="noopener"><code>bufferDeviceAddressMultiDevice</code></a> feature
  **must** be enabled

Valid Usage (Implicit)

- <a href="#VUID-vkCreateAccelerationStructureKHR-device-parameter"
  id="VUID-vkCreateAccelerationStructureKHR-device-parameter"></a>
  VUID-vkCreateAccelerationStructureKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateAccelerationStructureKHR-pCreateInfo-parameter"
  id="VUID-vkCreateAccelerationStructureKHR-pCreateInfo-parameter"></a>
  VUID-vkCreateAccelerationStructureKHR-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoKHR.html)
  structure

- <a href="#VUID-vkCreateAccelerationStructureKHR-pAllocator-parameter"
  id="VUID-vkCreateAccelerationStructureKHR-pAllocator-parameter"></a>
  VUID-vkCreateAccelerationStructureKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a
  href="#VUID-vkCreateAccelerationStructureKHR-pAccelerationStructure-parameter"
  id="VUID-vkCreateAccelerationStructureKHR-pAccelerationStructure-parameter"></a>
  VUID-vkCreateAccelerationStructureKHR-pAccelerationStructure-parameter  
  `pAccelerationStructure` **must** be a valid pointer to a
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoKHR.html),
[VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateAccelerationStructureKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
