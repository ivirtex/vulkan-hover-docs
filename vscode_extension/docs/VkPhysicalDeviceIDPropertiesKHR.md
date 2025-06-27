# VkPhysicalDeviceIDProperties(3) Manual Page

## Name

VkPhysicalDeviceIDProperties - Structure specifying IDs related to the
physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceIDProperties` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkPhysicalDeviceIDProperties {
    VkStructureType    sType;
    void*              pNext;
    uint8_t            deviceUUID[VK_UUID_SIZE];
    uint8_t            driverUUID[VK_UUID_SIZE];
    uint8_t            deviceLUID[VK_LUID_SIZE];
    uint32_t           deviceNodeMask;
    VkBool32           deviceLUIDValid;
} VkPhysicalDeviceIDProperties;
```

or the equivalent

``` c
// Provided by VK_KHR_external_fence_capabilities, VK_KHR_external_memory_capabilities, VK_KHR_external_semaphore_capabilities
typedef VkPhysicalDeviceIDProperties VkPhysicalDeviceIDPropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- `deviceUUID` is an array of `VK_UUID_SIZE` `uint8_t` values
  representing a universally unique identifier for the device.

- `driverUUID` is an array of `VK_UUID_SIZE` `uint8_t` values
  representing a universally unique identifier for the driver build in
  use by the device.

- `deviceLUID` is an array of `VK_LUID_SIZE` `uint8_t` values
  representing a locally unique identifier for the device.

- `deviceNodeMask` is a `uint32_t` bitfield identifying the node within
  a linked device adapter corresponding to the device.

- `deviceLUIDValid` is a boolean value that will be `VK_TRUE` if
  `deviceLUID` contains a valid LUID and `deviceNodeMask` contains a
  valid node mask, and `VK_FALSE` if they do not.

If the `VkPhysicalDeviceIDProperties` structure is included in the
`pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

`deviceUUID` **must** be immutable for a given device across instances,
processes, driver APIs, driver versions, and system reboots.

Applications **can** compare the `driverUUID` value across instance and
process boundaries, and **can** make similar queries in external APIs to
determine whether they are capable of sharing memory objects and
resources using them with the device.

`deviceUUID` and/or `driverUUID` **must** be used to determine whether a
particular external object can be shared between driver components,
where such a restriction exists as defined in the compatibility table
for the particular object type:

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#external-memory-handle-types-compatibility"
  target="_blank" rel="noopener">External memory handle types
  compatibility</a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#external-semaphore-handle-types-compatibility"
  target="_blank" rel="noopener">External semaphore handle types
  compatibility</a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#external-fence-handle-types-compatibility"
  target="_blank" rel="noopener">External fence handle types
  compatibility</a>

If `deviceLUIDValid` is `VK_FALSE`, the values of `deviceLUID` and
`deviceNodeMask` are undefined. If `deviceLUIDValid` is `VK_TRUE` and
Vulkan is running on the Windows operating system, the contents of
`deviceLUID` **can** be cast to an `LUID` object and **must** be equal
to the locally unique identifier of a `IDXGIAdapter1` object that
corresponds to `physicalDevice`. If `deviceLUIDValid` is `VK_TRUE`,
`deviceNodeMask` **must** contain exactly one bit. If Vulkan is running
on an operating system that supports the Direct3D 12 API and
`physicalDevice` corresponds to an individual device in a linked device
adapter, `deviceNodeMask` identifies the Direct3D 12 node corresponding
to `physicalDevice`. Otherwise, `deviceNodeMask` **must** be `1`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Although they have identical descriptions, <a
href="VkPhysicalDeviceIDProperties.html">VkPhysicalDeviceIDProperties</a>::<code>deviceUUID</code>
may differ from <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html">VkPhysicalDeviceProperties2</a>::<code>pipelineCacheUUID</code>.
The former is intended to identify and correlate devices across API and
driver boundaries, while the latter is used to identify a compatible
device and driver combination to use when serializing and de-serializing
pipeline state.</p>
<p>Implementations <strong>should</strong> return
<code>deviceUUID</code> values which are likely to be unique even in the
presence of multiple Vulkan implementations (such as a GPU driver and a
software renderer; two drivers for different GPUs; or the same Vulkan
driver running on two logically different devices).</p>
<p>Khronos' conformance testing is unable to guarantee that
<code>deviceUUID</code> values are actually unique, so implementors
<strong>should</strong> make their own best efforts to ensure this. In
particular, hard-coded <code>deviceUUID</code> values, especially
all-<code>0</code> bits, <strong>should</strong> never be used.</p>
<p>A combination of values unique to the vendor, the driver, and the
hardware environment can be used to provide a <code>deviceUUID</code>
which is unique to a high degree of certainty. Some possible inputs to
such a computation are:</p>
<ul>
<li><p>Information reported by <a
href="vkGetPhysicalDeviceProperties.html">vkGetPhysicalDeviceProperties</a></p></li>
<li><p>PCI device ID (if defined)</p></li>
<li><p>PCI bus ID, or similar system configuration information.</p></li>
<li><p>Driver binary checksums.</p></li>
</ul></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>While <a
href="VkPhysicalDeviceIDProperties.html">VkPhysicalDeviceIDProperties</a>::<code>deviceUUID</code>
is specified to remain consistent across driver versions and system
reboots, it is not intended to be usable as a serializable persistent
identifier for a device. It may change when a device is physically added
to, removed from, or moved to a different connector in a system while
that system is powered down. Further, there is no reasonable way to
verify with conformance testing that a given device retains the same
UUID in a given system across all driver versions supported in that
system. While implementations should make every effort to report
consistent device UUIDs across driver versions, applications should
avoid relying on the persistence of this value for uses other than
identifying compatible devices for external object sharing
purposes.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceIDProperties-sType-sType"
  id="VUID-VkPhysicalDeviceIDProperties-sType-sType"></a>
  VUID-VkPhysicalDeviceIDProperties-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceIDProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
