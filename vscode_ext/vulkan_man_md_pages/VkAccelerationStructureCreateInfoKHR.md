# VkAccelerationStructureCreateInfoKHR(3) Manual Page

## Name

VkAccelerationStructureCreateInfoKHR - Structure specifying the
parameters of a newly created acceleration structure object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAccelerationStructureCreateInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_acceleration_structure
typedef struct VkAccelerationStructureCreateInfoKHR {
    VkStructureType                          sType;
    const void*                              pNext;
    VkAccelerationStructureCreateFlagsKHR    createFlags;
    VkBuffer                                 buffer;
    VkDeviceSize                             offset;
    VkDeviceSize                             size;
    VkAccelerationStructureTypeKHR           type;
    VkDeviceAddress                          deviceAddress;
} VkAccelerationStructureCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `createFlags` is a bitmask of
  [VkAccelerationStructureCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateFlagBitsKHR.html)
  specifying additional creation parameters of the acceleration
  structure.

- `buffer` is the buffer on which the acceleration structure will be
  stored.

- `offset` is an offset in bytes from the base address of the buffer at
  which the acceleration structure will be stored, and **must** be a
  multiple of `256`.

- `size` is the size required for the acceleration structure.

- `type` is a
  [VkAccelerationStructureTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureTypeKHR.html)
  value specifying the type of acceleration structure that will be
  created.

- `deviceAddress` is the device address requested for the acceleration
  structure if the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-accelerationStructureCaptureReplay"
  target="_blank"
  rel="noopener"><code>accelerationStructureCaptureReplay</code></a>
  feature is being used. If `deviceAddress` is zero, no specific address
  is requested.

## <a href="#_description" class="anchor"></a>Description

Applications **should** avoid creating acceleration structures with
application-provided addresses and implementation-provided addresses in
the same process, to reduce the likelihood of
`VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR` errors.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>The expected usage for this is that a trace capture/replay tool will
add the <code>VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT</code>
flag to all buffers that use
<code>VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT</code>, and will add
<code>VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT</code> to all buffers
used as storage for an acceleration structure where
<code>deviceAddress</code> is not zero. This also means that the tool
will need to add <code>VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT</code> to
memory allocations to allow the flag to be set where the application may
not have otherwise required it. During capture the tool will save the
queried opaque device addresses in the trace. During replay, the buffers
will be created specifying the original address so any address values
stored in the trace data will remain valid.</p>
<p>Implementations are expected to separate such buffers in the GPU
address space so normal allocations will avoid using these addresses.
Apps/tools should avoid mixing app-provided and implementation-provided
addresses for buffers created with
<code>VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT</code>, to
avoid address space allocation conflicts.</p></td>
</tr>
</tbody>
</table>

Applications **should** create an acceleration structure with a specific
[VkAccelerationStructureTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureTypeKHR.html)
other than `VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p><code>VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR</code> is intended
to be used by API translation layers. This can be used at acceleration
structure creation time in cases where the actual acceleration structure
type (top or bottom) is not yet known. The actual acceleration structure
type must be specified as
<code>VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_KHR</code> or
<code>VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_KHR</code> when the
build is performed.</p></td>
</tr>
</tbody>
</table>

If the acceleration structure will be the target of a build operation,
the required size for an acceleration structure **can** be queried with
[vkGetAccelerationStructureBuildSizesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureBuildSizesKHR.html).
If the acceleration structure is going to be the target of a compacting
copy,
[vkCmdWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteAccelerationStructuresPropertiesKHR.html)
or
[vkWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkWriteAccelerationStructuresPropertiesKHR.html)
**can** be used to obtain the compacted size required.

If the acceleration structure will be the target of a build operation
with `VK_BUILD_ACCELERATION_STRUCTURE_MOTION_BIT_NV` it **must** include
`VK_ACCELERATION_STRUCTURE_CREATE_MOTION_BIT_NV` in `createFlags` and
include
[VkAccelerationStructureMotionInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMotionInfoNV.html)
as an extension structure in `pNext` with the number of instances as
metadata for the object.

Valid Usage

- <a href="#VUID-VkAccelerationStructureCreateInfoKHR-deviceAddress-03612"
  id="VUID-VkAccelerationStructureCreateInfoKHR-deviceAddress-03612"></a>
  VUID-VkAccelerationStructureCreateInfoKHR-deviceAddress-03612  
  If `deviceAddress` is not zero, `createFlags` **must** include
  `VK_ACCELERATION_STRUCTURE_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_KHR`

- <a href="#VUID-VkAccelerationStructureCreateInfoKHR-deviceAddress-09488"
  id="VUID-VkAccelerationStructureCreateInfoKHR-deviceAddress-09488"></a>
  VUID-VkAccelerationStructureCreateInfoKHR-deviceAddress-09488  
  If `deviceAddress` is not zero, it **must** have been retrieved from
  an identically created acceleration structure, except for `buffer` and
  `deviceAddress`

- <a href="#VUID-VkAccelerationStructureCreateInfoKHR-deviceAddress-09489"
  id="VUID-VkAccelerationStructureCreateInfoKHR-deviceAddress-09489"></a>
  VUID-VkAccelerationStructureCreateInfoKHR-deviceAddress-09489  
  If `deviceAddress` is not zero, `buffer` **must** have been created
  identically to the `buffer` used to create the acceleration structure
  from which `deviceAddress` was retrieved, except for
  [VkBufferOpaqueCaptureAddressCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferOpaqueCaptureAddressCreateInfo.html)::`opaqueCaptureAddress`

- <a href="#VUID-VkAccelerationStructureCreateInfoKHR-deviceAddress-09490"
  id="VUID-VkAccelerationStructureCreateInfoKHR-deviceAddress-09490"></a>
  VUID-VkAccelerationStructureCreateInfoKHR-deviceAddress-09490  
  If `deviceAddress` is not zero, `buffer` **must** have been created
  with a
  [VkBufferOpaqueCaptureAddressCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferOpaqueCaptureAddressCreateInfo.html)::`opaqueCaptureAddress`
  that was retrieved from
  [vkGetBufferOpaqueCaptureAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferOpaqueCaptureAddress.html)
  for the `buffer` that was used to create the acceleration structure
  from which `deviceAddress` was retrieved

- <a href="#VUID-VkAccelerationStructureCreateInfoKHR-createFlags-03613"
  id="VUID-VkAccelerationStructureCreateInfoKHR-createFlags-03613"></a>
  VUID-VkAccelerationStructureCreateInfoKHR-createFlags-03613  
  If `createFlags` includes
  `VK_ACCELERATION_STRUCTURE_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_KHR`,
  [VkPhysicalDeviceAccelerationStructureFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceAccelerationStructureFeaturesKHR.html)::`accelerationStructureCaptureReplay`
  **must** be `VK_TRUE`

- <a href="#VUID-VkAccelerationStructureCreateInfoKHR-buffer-03614"
  id="VUID-VkAccelerationStructureCreateInfoKHR-buffer-03614"></a>
  VUID-VkAccelerationStructureCreateInfoKHR-buffer-03614  
  `buffer` **must** have been created with a `usage` value containing
  `VK_BUFFER_USAGE_ACCELERATION_STRUCTURE_STORAGE_BIT_KHR`

- <a href="#VUID-VkAccelerationStructureCreateInfoKHR-buffer-03615"
  id="VUID-VkAccelerationStructureCreateInfoKHR-buffer-03615"></a>
  VUID-VkAccelerationStructureCreateInfoKHR-buffer-03615  
  `buffer` **must** not have been created with
  `VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT`

- <a href="#VUID-VkAccelerationStructureCreateInfoKHR-offset-03616"
  id="VUID-VkAccelerationStructureCreateInfoKHR-offset-03616"></a>
  VUID-VkAccelerationStructureCreateInfoKHR-offset-03616  
  The sum of `offset` and `size` **must** be less than the size of
  `buffer`

- <a href="#VUID-VkAccelerationStructureCreateInfoKHR-offset-03734"
  id="VUID-VkAccelerationStructureCreateInfoKHR-offset-03734"></a>
  VUID-VkAccelerationStructureCreateInfoKHR-offset-03734  
  `offset` **must** be a multiple of `256` bytes

- <a href="#VUID-VkAccelerationStructureCreateInfoKHR-createFlags-04954"
  id="VUID-VkAccelerationStructureCreateInfoKHR-createFlags-04954"></a>
  VUID-VkAccelerationStructureCreateInfoKHR-createFlags-04954  
  If `VK_ACCELERATION_STRUCTURE_CREATE_MOTION_BIT_NV` is set in
  `createFlags` and `type` is
  `VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_KHR`, one member of the
  `pNext` chain **must** be a pointer to a valid instance of
  [VkAccelerationStructureMotionInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMotionInfoNV.html)

- <a href="#VUID-VkAccelerationStructureCreateInfoKHR-createFlags-04955"
  id="VUID-VkAccelerationStructureCreateInfoKHR-createFlags-04955"></a>
  VUID-VkAccelerationStructureCreateInfoKHR-createFlags-04955  
  If any geometry includes
  `VkAccelerationStructureGeometryMotionTrianglesDataNV` then
  `createFlags` **must** contain
  `VK_ACCELERATION_STRUCTURE_CREATE_MOTION_BIT_NV`

- <a href="#VUID-VkAccelerationStructureCreateInfoKHR-createFlags-08108"
  id="VUID-VkAccelerationStructureCreateInfoKHR-createFlags-08108"></a>
  VUID-VkAccelerationStructureCreateInfoKHR-createFlags-08108  
  If `createFlags` includes
  `VK_ACCELERATION_STRUCTURE_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`,
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-descriptorBufferCaptureReplay"
  target="_blank"
  rel="noopener"><code>descriptorBufferCaptureReplay</code></a> feature
  **must** be enabled

- <a href="#VUID-VkAccelerationStructureCreateInfoKHR-pNext-08109"
  id="VUID-VkAccelerationStructureCreateInfoKHR-pNext-08109"></a>
  VUID-VkAccelerationStructureCreateInfoKHR-pNext-08109  
  If the `pNext` chain includes a
  [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html)
  structure, `createFlags` **must** contain
  `VK_ACCELERATION_STRUCTURE_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`

Valid Usage (Implicit)

- <a href="#VUID-VkAccelerationStructureCreateInfoKHR-sType-sType"
  id="VUID-VkAccelerationStructureCreateInfoKHR-sType-sType"></a>
  VUID-VkAccelerationStructureCreateInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_CREATE_INFO_KHR`

- <a href="#VUID-VkAccelerationStructureCreateInfoKHR-pNext-pNext"
  id="VUID-VkAccelerationStructureCreateInfoKHR-pNext-pNext"></a>
  VUID-VkAccelerationStructureCreateInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkAccelerationStructureMotionInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMotionInfoNV.html)
  or
  [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html)

- <a href="#VUID-VkAccelerationStructureCreateInfoKHR-sType-unique"
  id="VUID-VkAccelerationStructureCreateInfoKHR-sType-unique"></a>
  VUID-VkAccelerationStructureCreateInfoKHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a
  href="#VUID-VkAccelerationStructureCreateInfoKHR-createFlags-parameter"
  id="VUID-VkAccelerationStructureCreateInfoKHR-createFlags-parameter"></a>
  VUID-VkAccelerationStructureCreateInfoKHR-createFlags-parameter  
  `createFlags` **must** be a valid combination of
  [VkAccelerationStructureCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateFlagBitsKHR.html)
  values

- <a href="#VUID-VkAccelerationStructureCreateInfoKHR-buffer-parameter"
  id="VUID-VkAccelerationStructureCreateInfoKHR-buffer-parameter"></a>
  VUID-VkAccelerationStructureCreateInfoKHR-buffer-parameter  
  `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-VkAccelerationStructureCreateInfoKHR-type-parameter"
  id="VUID-VkAccelerationStructureCreateInfoKHR-type-parameter"></a>
  VUID-VkAccelerationStructureCreateInfoKHR-type-parameter  
  `type` **must** be a valid
  [VkAccelerationStructureTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureTypeKHR.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkAccelerationStructureCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateFlagsKHR.html),
[VkAccelerationStructureTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureTypeKHR.html),
[VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateAccelerationStructureKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAccelerationStructureCreateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
