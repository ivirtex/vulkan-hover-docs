# VK\_KHR\_external\_memory\_capabilities(3) Manual Page

## Name

VK\_KHR\_external\_memory\_capabilities - instance extension



## [](#_registered_extension_number)Registered Extension Number

72

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.1](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.1-promotions)

## [](#_contact)Contact

- James Jones [\[GitHub\]cubanismo](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_external_memory_capabilities%5D%20%40cubanismo%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_external_memory_capabilities%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2016-10-17

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- Interacts with `VK_KHR_dedicated_allocation`.
- Interacts with `VK_NV_dedicated_allocation`.

**Contributors**

- Ian Elliott, Google
- Jesse Hall, Google
- James Jones, NVIDIA

## [](#_description)Description

An application may wish to reference device memory in multiple Vulkan logical devices or instances, in multiple processes, and/or in multiple APIs. This extension provides a set of capability queries and handle definitions that allow an application to determine what types of “external” memory handles an implementation supports for a given set of use cases.

## [](#_promotion_to_vulkan_1_1)Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with the KHR suffix omitted. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_commands)New Commands

- [vkGetPhysicalDeviceExternalBufferPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceExternalBufferPropertiesKHR.html)

## [](#_new_structures)New Structures

- [VkExternalBufferPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalBufferPropertiesKHR.html)
- [VkExternalMemoryPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryPropertiesKHR.html)
- [VkPhysicalDeviceExternalBufferInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalBufferInfoKHR.html)
- Extending [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatProperties2.html):
  
  - [VkExternalImageFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalImageFormatPropertiesKHR.html)
- Extending [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html):
  
  - [VkPhysicalDeviceExternalImageFormatInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalImageFormatInfoKHR.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceIDPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceIDPropertiesKHR.html)

## [](#_new_enums)New Enums

- [VkExternalMemoryFeatureFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryFeatureFlagBitsKHR.html)
- [VkExternalMemoryHandleTypeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBitsKHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkExternalMemoryFeatureFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryFeatureFlagsKHR.html)
- [VkExternalMemoryHandleTypeFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_EXTERNAL_MEMORY_CAPABILITIES_EXTENSION_NAME`
- `VK_KHR_EXTERNAL_MEMORY_CAPABILITIES_SPEC_VERSION`
- `VK_LUID_SIZE_KHR`
- Extending [VkExternalMemoryFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryFeatureFlagBits.html):
  
  - `VK_EXTERNAL_MEMORY_FEATURE_DEDICATED_ONLY_BIT_KHR`
  - `VK_EXTERNAL_MEMORY_FEATURE_EXPORTABLE_BIT_KHR`
  - `VK_EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT_KHR`
- Extending [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html):
  
  - `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT_KHR`
  - `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT_KHR`
  - `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT_KHR`
  - `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_RESOURCE_BIT_KHR`
  - `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT_KHR`
  - `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR`
  - `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES_KHR`
  - `VK_STRUCTURE_TYPE_EXTERNAL_IMAGE_FORMAT_PROPERTIES_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_BUFFER_INFO_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_IMAGE_FORMAT_INFO_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES_KHR`

## [](#_issues)Issues

1\) Why do so many external memory capabilities need to be queried on a per-memory-handle-type basis?

**PROPOSED RESOLUTION**: This is because some handle types are based on OS-native objects that have far more limited capabilities than the very generic Vulkan memory objects. Not all memory handle types can name memory objects that support 3D images, for example. Some handle types cannot even support the deferred image and memory binding behavior of Vulkan and require specifying the image when allocating or importing the memory object.

2\) Do the [VkExternalImageFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalImageFormatPropertiesKHR.html) and [VkExternalBufferPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalBufferPropertiesKHR.html) structs need to include a list of memory type bits that support the given handle type?

**PROPOSED RESOLUTION**: No. The memory types that do not support the handle types will simply be filtered out of the results returned by [vkGetImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageMemoryRequirements.html) and [vkGetBufferMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferMemoryRequirements.html) when a set of handle types was specified at image or buffer creation time.

3\) Should the non-opaque handle types be moved to their own extension?

**PROPOSED RESOLUTION**: Perhaps. However, defining the handle type bits does very little and does not require any platform-specific types on its own, and it is easier to maintain the bitfield values in a single extension for now. Presumably more handle types could be added by separate extensions though, and it would be midly weird to have some platform-specific ones defined in the core spec and some in extensions

4\) Do we need a `D3D11_TILEPOOL` type?

**PROPOSED RESOLUTION**: No. This is technically possible, but the synchronization is awkward. D3D11 surfaces must be synchronized using shared mutexes, and these synchronization primitives are shared by the entire memory object, so D3D11 shared allocations divided among multiple buffer and image bindings may be difficult to synchronize.

5\) Should the Windows 7-compatible handle types be named “KMT” handles or “GLOBAL\_SHARE” handles?

**PROPOSED RESOLUTION**: KMT, simply because it is more concise.

6\) How do applications identify compatible devices and drivers across instance, process, and API boundaries when sharing memory?

**PROPOSED RESOLUTION**: New device properties are exposed that allow applications to correctly correlate devices and drivers. A device and driver UUID that must both match to ensure sharing compatibility between two Vulkan instances, or a Vulkan instance and an extensible external API are added. To allow correlating with Direct3D devices, a device LUID is added that corresponds to a DXGI adapter LUID. A driver ID is not needed for Direct3D because mismatched driver component versions are not currently supported on the Windows OS. Should support for such configurations be introduced at the OS level, further Vulkan extensions would be needed to correlate userspace component builds.

## [](#_version_history)Version History

- Revision 1, 2016-10-17 (James Jones)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_external_memory_capabilities)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0