# VkPhysicalDeviceProperties(3) Manual Page

## Name

VkPhysicalDeviceProperties - Structure specifying physical device properties



## [](#_c_specification)C Specification

The `VkPhysicalDeviceProperties` structure is defined as:

```c++
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

## [](#_members)Members

- `apiVersion` is the version of Vulkan supported by the device, encoded as described in [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-coreversions-versionnumbers](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-coreversions-versionnumbers).
- `driverVersion` is the vendor-specified version of the driver.
- `vendorID` is a unique identifier for the *vendor* (see below) of the physical device.
- `deviceID` is a unique identifier for the physical device among devices available from the vendor.
- `deviceType` is a [VkPhysicalDeviceType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceType.html) specifying the type of device.
- `deviceName` is an array of `VK_MAX_PHYSICAL_DEVICE_NAME_SIZE` `char` containing a null-terminated UTF-8 string which is the name of the device.
- `pipelineCacheUUID` is an array of `VK_UUID_SIZE` `uint8_t` values representing a universally unique identifier for the device.
- `limits` is the [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLimits.html) structure specifying device-specific limits of the physical device. See [Limits](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits) for details.
- `sparseProperties` is the [VkPhysicalDeviceSparseProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSparseProperties.html) structure specifying various sparse related properties of the physical device. See [Sparse Properties](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#sparsememory-physicalprops) for details.

## [](#_description)Description

Note

The value of `apiVersion` **may** be different than the version returned by [vkEnumerateInstanceVersion](https://registry.khronos.org/vulkan/specs/latest/man/html/vkEnumerateInstanceVersion.html); either higher or lower. In such cases, the application **must** not use functionality that exceeds the version of Vulkan associated with a given object. The `pApiVersion` parameter returned by [vkEnumerateInstanceVersion](https://registry.khronos.org/vulkan/specs/latest/man/html/vkEnumerateInstanceVersion.html) is the version associated with a [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) and its children, except for a [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) and its children. `VkPhysicalDeviceProperties`::`apiVersion` is the version associated with a [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) and its children.

Note

The encoding of `driverVersion` is implementation-defined. It **may** not use the same encoding as `apiVersion`. Applications should follow information from the *vendor* on how to extract the version information from `driverVersion`.

On implementations that claim support for the [Roadmap 2022](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#roadmap-2022) profile, the major and minor version expressed by `apiVersion` **must** be at least Vulkan 1.3.

The `vendorID` and `deviceID` fields are provided to allow applications to adapt to device characteristics that are not adequately exposed by other Vulkan queries.

Note

These **may** include performance profiles, hardware errata, or other characteristics.

The *vendor* identified by `vendorID` is the entity responsible for the most salient characteristics of the underlying implementation of the [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) being queried.

Note

For example, in the case of a discrete GPU implementation, this **should** be the GPU chipset vendor. In the case of a hardware accelerator integrated into a system-on-chip (SoC), this **should** be the supplier of the silicon IP used to create the accelerator.

If the vendor has a [PCI vendor ID](https://pcisig.com/membership/member-companies), the low 16 bits of `vendorID` **must** contain that PCI vendor ID, and the remaining bits **must** be zero. Otherwise, the value returned **must** be a valid Khronos vendor ID, obtained as described in the [Vulkan Documentation and Extensions: Procedures and Conventions](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vulkan-styleguide) document in the section “Registering a Vendor ID with Khronos”. Khronos vendor IDs are allocated starting at 0x10000, to distinguish them from the PCI vendor ID namespace. Khronos vendor IDs are symbolically defined in the [VkVendorId](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVendorId.html) type.

The vendor is also responsible for the value returned in `deviceID`. If the implementation is driven primarily by a [PCI device](https://pcisig.com/) with a [PCI device ID](https://pcisig.com/), the low 16 bits of `deviceID` **must** contain that PCI device ID, and the remaining bits **must** be zero. Otherwise, the choice of what values to return **may** be dictated by operating system or platform policies - but **should** uniquely identify both the device version and any major configuration options (for example, core count in the case of multicore devices).

Note

The same device ID **should** be used for all physical implementations of that device version and configuration. For example, all uses of a specific silicon IP GPU version and configuration **should** use the same device ID, even if those uses occur in different SoCs.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLimits.html), [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html), [VkPhysicalDeviceSparseProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSparseProperties.html), [VkPhysicalDeviceType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceType.html), [vkGetPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceProperties)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0