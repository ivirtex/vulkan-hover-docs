# VkPhysicalDeviceProperties(3) Manual Page

## Name

VkPhysicalDeviceProperties - Structure specifying physical device
properties



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceProperties` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkPhysicalDeviceProperties {
    uint32_t                            apiVersion;
    uint32_t                            driverVersion;
    uint32_t                            vendorID;
    uint32_t                            deviceID;
    VkPhysicalDeviceType                deviceType;
    char                                deviceName[VK_MAX_PHYSICAL_DEVICE_NAME_SIZE];
    uint8_t                             pipelineCacheUUID[VK_UUID_SIZE];
    VkPhysicalDeviceLimits              limits;
    VkPhysicalDeviceSparseProperties    sparseProperties;
} VkPhysicalDeviceProperties;
```

## <a href="#_members" class="anchor"></a>Members

- `apiVersion` is the version of Vulkan supported by the device, encoded
  as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-coreversions-versionnumbers"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-coreversions-versionnumbers</a>.

- `driverVersion` is the vendor-specified version of the driver.

- `vendorID` is a unique identifier for the *vendor* (see below) of the
  physical device.

- `deviceID` is a unique identifier for the physical device among
  devices available from the vendor.

- `deviceType` is a [VkPhysicalDeviceType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceType.html)
  specifying the type of device.

- `deviceName` is an array of `VK_MAX_PHYSICAL_DEVICE_NAME_SIZE` `char`
  containing a null-terminated UTF-8 string which is the name of the
  device.

- `pipelineCacheUUID` is an array of `VK_UUID_SIZE` `uint8_t` values
  representing a universally unique identifier for the device.

- `limits` is the [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)
  structure specifying device-specific limits of the physical device.
  See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits"
  target="_blank" rel="noopener">Limits</a> for details.

- `sparseProperties` is the
  [VkPhysicalDeviceSparseProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSparseProperties.html)
  structure specifying various sparse related properties of the physical
  device. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#sparsememory-physicalprops"
  target="_blank" rel="noopener">Sparse Properties</a> for details.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The value of <code>apiVersion</code> <strong>may</strong> be
different than the version returned by <a
href="vkEnumerateInstanceVersion.html">vkEnumerateInstanceVersion</a>;
either higher or lower. In such cases, the application
<strong>must</strong> not use functionality that exceeds the version of
Vulkan associated with a given object. The <code>pApiVersion</code>
parameter returned by <a
href="vkEnumerateInstanceVersion.html">vkEnumerateInstanceVersion</a> is
the version associated with a <a href="VkInstance.html">VkInstance</a>
and its children, except for a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html">VkPhysicalDevice</a> and its children.
<code>VkPhysicalDeviceProperties</code>::<code>apiVersion</code> is the
version associated with a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html">VkPhysicalDevice</a> and its
children.</p></td>
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
<p>The encoding of <code>driverVersion</code> is implementation-defined.
It <strong>may</strong> not use the same encoding as
<code>apiVersion</code>. Applications should follow information from the
<em>vendor</em> on how to extract the version information from
<code>driverVersion</code>.</p></td>
</tr>
</tbody>
</table>

On implementations that claim support for the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#roadmap-2022"
target="_blank" rel="noopener">Roadmap 2022</a> profile, the major and
minor version expressed by `apiVersion` **must** be at least Vulkan 1.3.

The `vendorID` and `deviceID` fields are provided to allow applications
to adapt to device characteristics that are not adequately exposed by
other Vulkan queries.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>These <strong>may</strong> include performance profiles, hardware
errata, or other characteristics.</p></td>
</tr>
</tbody>
</table>

The *vendor* identified by `vendorID` is the entity responsible for the
most salient characteristics of the underlying implementation of the
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) being queried.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>For example, in the case of a discrete GPU implementation, this
<strong>should</strong> be the GPU chipset vendor. In the case of a
hardware accelerator integrated into a system-on-chip (SoC), this
<strong>should</strong> be the supplier of the silicon IP used to create
the accelerator.</p></td>
</tr>
</tbody>
</table>

If the vendor has a [PCI vendor
ID](https://pcisig.com/membership/member-companies), the low 16 bits of
`vendorID` **must** contain that PCI vendor ID, and the remaining bits
**must** be set to zero. Otherwise, the value returned **must** be a
valid Khronos vendor ID, obtained as described in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vulkan-styleguide"
target="_blank" rel="noopener">Vulkan Documentation and Extensions:
Procedures and Conventions</a> document in the section “Registering a
Vendor ID with Khronos”. Khronos vendor IDs are allocated starting at
0x10000, to distinguish them from the PCI vendor ID namespace. Khronos
vendor IDs are symbolically defined in the [VkVendorId](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVendorId.html)
type.

The vendor is also responsible for the value returned in `deviceID`. If
the implementation is driven primarily by a [PCI
device](https://pcisig.com/) with a [PCI device
ID](https://pcisig.com/), the low 16 bits of `deviceID` **must** contain
that PCI device ID, and the remaining bits **must** be set to zero.
Otherwise, the choice of what values to return **may** be dictated by
operating system or platform policies - but **should** uniquely identify
both the device version and any major configuration options (for
example, core count in the case of multicore devices).

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The same device ID <strong>should</strong> be used for all physical
implementations of that device version and configuration. For example,
all uses of a specific silicon IP GPU version and configuration
<strong>should</strong> use the same device ID, even if those uses occur
in different SoCs.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html),
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html),
[VkPhysicalDeviceSparseProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSparseProperties.html),
[VkPhysicalDeviceType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceType.html),
[vkGetPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
