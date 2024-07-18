# VK_NV_external_memory(3) Manual Page

## Name

VK_NV_external_memory - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

57

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_NV_external_memory_capabilities](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_external_memory_capabilities.html)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Deprecated* by [VK_KHR_external_memory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory.html)
  extension

  - Which in turn was *promoted* to <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.1-promotions"
    target="_blank" rel="noopener">Vulkan 1.1</a>

## <a href="#_contact" class="anchor"></a>Contact

- James Jones <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_external_memory%5D%20@cubanismo%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_external_memory%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>cubanismo</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2016-08-19

**IP Status**  
No known IP claims.

**Contributors**  
- James Jones, NVIDIA

- Carsten Rohde, NVIDIA

## <a href="#_description" class="anchor"></a>Description

Applications may wish to export memory to other Vulkan instances or
other APIs, or import memory from other Vulkan instances or other APIs
to enable Vulkan workloads to be split up across application module,
process, or API boundaries. This extension enables applications to
create exportable Vulkan memory objects such that the underlying
resources can be referenced outside the Vulkan instance that created
them.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html):

  - [VkExternalMemoryImageCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfoNV.html)

- Extending [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html):

  - [VkExportMemoryAllocateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfoNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_EXTERNAL_MEMORY_EXTENSION_NAME`

- `VK_NV_EXTERNAL_MEMORY_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO_NV`

  - `VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO_NV`

## <a href="#_issues" class="anchor"></a>Issues

1\) If memory objects are shared between processes and APIs, is this
considered aliasing according to the rules outlined in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-memory-aliasing"
target="_blank" rel="noopener">Memory Aliasing</a> section?

**RESOLVED**: Yes, but strict exceptions to the rules are added to allow
some forms of aliasing in these cases. Further, other extensions may
build upon these new aliasing rules to define specific support usage
within Vulkan for imported native memory objects, or memory objects from
other APIs.

2\) Are new image layouts or metadata required to specify image layouts
and layout transitions compatible with non-Vulkan APIs, or with other
instances of the same Vulkan driver?

**RESOLVED**: No. Separate instances of the same Vulkan driver running
on the same GPU should have identical internal layout semantics, so
applications have the tools they need to ensure views of images are
consistent between the two instances. Other APIs will fall into two
categories: Those that are Vulkan compatible (a term to be defined by
subsequent interopability extensions), or Vulkan incompatible. When
sharing images with Vulkan incompatible APIs, the Vulkan image must be
transitioned to the `VK_IMAGE_LAYOUT_GENERAL` layout before handing it
off to the external API.

Note this does not attempt to address cross-device transitions, nor
transitions to engines on the same device which are not visible within
the Vulkan API. Both of these are beyond the scope of this extension.

## <a href="#_examples" class="anchor"></a>Examples

``` c
    // TODO: Write some sample code here.
```

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-08-19 (James Jones)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_external_memory"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
