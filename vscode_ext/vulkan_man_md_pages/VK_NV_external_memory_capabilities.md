# VK_NV_external_memory_capabilities(3) Manual Page

## Name

VK_NV_external_memory_capabilities - instance extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

56

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Deprecated* by
  [VK_KHR_external_memory_capabilities](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory_capabilities.html)
  extension

  - Which in turn was *promoted* to <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.1-promotions"
    target="_blank" rel="noopener">Vulkan 1.1</a>

## <a href="#_contact" class="anchor"></a>Contact

- James Jones <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_external_memory_capabilities%5D%20@cubanismo%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_external_memory_capabilities%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>cubanismo</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2016-08-19

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- Interacts with Vulkan 1.1.

- Interacts with
  [`VK_KHR_dedicated_allocation`](VK_KHR_dedicated_allocation.html).

- Interacts with
  [`VK_NV_dedicated_allocation`](VK_NV_dedicated_allocation.html).

**Contributors**  
- James Jones, NVIDIA

## <a href="#_description" class="anchor"></a>Description

Applications may wish to import memory from the Direct 3D API, or export
memory to other Vulkan instances. This extension provides a set of
capability queries that allow applications determine what types of win32
memory handles an implementation supports for a given set of use cases.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetPhysicalDeviceExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalImageFormatPropertiesNV.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalImageFormatPropertiesNV.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkExternalMemoryFeatureFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryFeatureFlagBitsNV.html)

- [VkExternalMemoryHandleTypeFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBitsNV.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkExternalMemoryFeatureFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryFeatureFlagsNV.html)

- [VkExternalMemoryHandleTypeFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagsNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_EXTERNAL_MEMORY_CAPABILITIES_EXTENSION_NAME`

- `VK_NV_EXTERNAL_MEMORY_CAPABILITIES_SPEC_VERSION`

## <a href="#_issues" class="anchor"></a>Issues

1\) Why do so many external memory capabilities need to be queried on a
per-memory-handle-type basis?

**RESOLVED**: This is because some handle types are based on OS-native
objects that have far more limited capabilities than the very generic
Vulkan memory objects. Not all memory handle types can name memory
objects that support 3D images, for example. Some handle types cannot
even support the deferred image and memory binding behavior of Vulkan
and require specifying the image when allocating or importing the memory
object.

2\) Does the
[VkExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalImageFormatPropertiesNV.html)
struct need to include a list of memory type bits that support the given
handle type?

**RESOLVED**: No. The memory types that do not support the handle types
will simply be filtered out of the results returned by
[vkGetImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements.html) when a
set of handle types was specified at image creation time.

3\) Should the non-opaque handle types be moved to their own extension?

**RESOLVED**: Perhaps. However, defining the handle type bits does very
little and does not require any platform-specific types on its own, and
it is easier to maintain the bitmask values in a single extension for
now. Presumably more handle types could be added by separate extensions
though, and it would be midly weird to have some platform-specific ones
defined in the core spec and some in extensions

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-08-19 (James Jones)

  - Initial version

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_external_memory_capabilities"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
