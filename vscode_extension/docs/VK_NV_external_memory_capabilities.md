# VK\_NV\_external\_memory\_capabilities(3) Manual Page

## Name

VK\_NV\_external\_memory\_capabilities - instance extension



## [](#_registered_extension_number)Registered Extension Number

56

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_deprecation_state)Deprecation State

- *Deprecated* by [VK\_KHR\_external\_memory\_capabilities](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory_capabilities.html) extension
  
  - Which in turn was *promoted* to [Vulkan 1.1](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.1-promotions)

## [](#_contact)Contact

- James Jones [\[GitHub\]cubanismo](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_external_memory_capabilities%5D%20%40cubanismo%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_external_memory_capabilities%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2016-08-19

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- Interacts with Vulkan 1.1.
- Interacts with `VK_KHR_dedicated_allocation`.
- Interacts with `VK_NV_dedicated_allocation`.

**Contributors**

- James Jones, NVIDIA

## [](#_description)Description

Applications may wish to import memory from the Direct 3D API, or export memory to other Vulkan instances. This extension provides a set of capability queries that allow applications determine what types of win32 memory handles an implementation supports for a given set of use cases.

## [](#_new_commands)New Commands

- [vkGetPhysicalDeviceExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceExternalImageFormatPropertiesNV.html)

## [](#_new_structures)New Structures

- [VkExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalImageFormatPropertiesNV.html)

## [](#_new_enums)New Enums

- [VkExternalMemoryFeatureFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryFeatureFlagBitsNV.html)
- [VkExternalMemoryHandleTypeFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBitsNV.html)

## [](#_new_bitmasks)New Bitmasks

- [VkExternalMemoryFeatureFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryFeatureFlagsNV.html)
- [VkExternalMemoryHandleTypeFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagsNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_EXTERNAL_MEMORY_CAPABILITIES_EXTENSION_NAME`
- `VK_NV_EXTERNAL_MEMORY_CAPABILITIES_SPEC_VERSION`

## [](#_issues)Issues

1\) Why do so many external memory capabilities need to be queried on a per-memory-handle-type basis?

**RESOLVED**: This is because some handle types are based on OS-native objects that have far more limited capabilities than the very generic Vulkan memory objects. Not all memory handle types can name memory objects that support 3D images, for example. Some handle types cannot even support the deferred image and memory binding behavior of Vulkan and require specifying the image when allocating or importing the memory object.

2\) Does the [VkExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalImageFormatPropertiesNV.html) structure need to include a list of memory type bits that support the given handle type?

**RESOLVED**: No. The memory types that do not support the handle types will simply be filtered out of the results returned by [vkGetImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageMemoryRequirements.html) when a set of handle types was specified at image creation time.

3\) Should the non-opaque handle types be moved to their own extension?

**RESOLVED**: Perhaps. However, defining the handle type bits does very little and does not require any platform-specific types on its own, and it is easier to maintain the bitmask values in a single extension for now. Presumably more handle types could be added by separate extensions though, and it would be midly weird to have some platform-specific ones defined in the core spec and some in extensions

## [](#_version_history)Version History

- Revision 1, 2016-08-19 (James Jones)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_external_memory_capabilities).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0