# VK\_NV\_external\_memory(3) Manual Page

## Name

VK\_NV\_external\_memory - device extension



## [](#_registered_extension_number)Registered Extension Number

57

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_NV\_external\_memory\_capabilities](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_external_memory_capabilities.html)

## [](#_deprecation_state)Deprecation State

- *Deprecated* by [VK\_KHR\_external\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory.html) extension
  
  - Which in turn was *promoted* to [Vulkan 1.1](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.1-promotions)

## [](#_contact)Contact

- James Jones [\[GitHub\]cubanismo](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_external_memory%5D%20%40cubanismo%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_external_memory%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2016-08-19

**IP Status**

No known IP claims.

**Contributors**

- James Jones, NVIDIA
- Carsten Rohde, NVIDIA

## [](#_description)Description

Applications may wish to export memory to other Vulkan instances or other APIs, or import memory from other Vulkan instances or other APIs to enable Vulkan workloads to be split up across application module, process, or API boundaries. This extension enables applications to create exportable Vulkan memory objects such that the underlying resources can be referenced outside the Vulkan instance that created them.

## [](#_new_structures)New Structures

- Extending [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html):
  
  - [VkExternalMemoryImageCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryImageCreateInfoNV.html)
- Extending [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html):
  
  - [VkExportMemoryAllocateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMemoryAllocateInfoNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_EXTERNAL_MEMORY_EXTENSION_NAME`
- `VK_NV_EXTERNAL_MEMORY_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO_NV`
  - `VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO_NV`

## [](#_issues)Issues

1\) If memory objects are shared between processes and APIs, is this considered aliasing according to the rules outlined in the [Memory Aliasing](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-memory-aliasing) section?

**RESOLVED**: Yes, but strict exceptions to the rules are added to allow some forms of aliasing in these cases. Further, other extensions may build upon these new aliasing rules to define specific support usage within Vulkan for imported native memory objects, or memory objects from other APIs.

2\) Are new image layouts or metadata required to specify image layouts and layout transitions compatible with non-Vulkan APIs, or with other instances of the same Vulkan driver?

**RESOLVED**: No. Separate instances of the same Vulkan driver running on the same GPU should have identical internal layout semantics, so applications have the tools they need to ensure views of images are consistent between the two instances. Other APIs will fall into two categories: Those that are Vulkan compatible (a term to be defined by subsequent interopability extensions), or Vulkan incompatible. When sharing images with Vulkan incompatible APIs, the Vulkan image must be transitioned to the `VK_IMAGE_LAYOUT_GENERAL` layout before handing it off to the external API.

Note this does not attempt to address cross-device transitions, nor transitions to engines on the same device which are not visible within the Vulkan API. Both of these are beyond the scope of this extension.

## [](#_examples)Examples

```c++
    // TODO: Write some sample code here.
```

## [](#_version_history)Version History

- Revision 1, 2016-08-19 (James Jones)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_external_memory)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0