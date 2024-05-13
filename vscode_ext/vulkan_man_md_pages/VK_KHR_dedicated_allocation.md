# VK_KHR_dedicated_allocation(3) Manual Page

## Name

VK_KHR_dedicated_allocation - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

128

## <a href="#_revision" class="anchor"></a>Revision

3

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_memory_requirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_memory_requirements2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.1-promotions"
  target="_blank" rel="noopener">Vulkan 1.1</a>

## <a href="#_contact" class="anchor"></a>Contact

- James Jones <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_dedicated_allocation%5D%20@cubanismo%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_dedicated_allocation%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>cubanismo</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-09-05

**IP Status**  
No known IP claims.

**Contributors**  
- Jeff Bolz, NVIDIA

- Faith Ekstrand, Intel

## <a href="#_description" class="anchor"></a>Description

This extension enables resources to be bound to a dedicated allocation,
rather than suballocated. For any particular resource, applications
**can** query whether a dedicated allocation is recommended, in which
case using a dedicated allocation **may** improve the performance of
access to that resource. Normal device memory allocations must support
multiple resources per allocation, memory aliasing and sparse binding,
which could interfere with some optimizations. Applications should query
the implementation for when a dedicated allocation **may** be beneficial
by adding a `VkMemoryDedicatedRequirementsKHR` structure to the `pNext`
chain of the `VkMemoryRequirements2` structure passed as the
`pMemoryRequirements` parameter of a call to
`vkGetBufferMemoryRequirements2` or `vkGetImageMemoryRequirements2`.
Certain external handle types and external images or buffers **may**
also depend on dedicated allocations on implementations that associate
image or buffer metadata with OS-level memory objects.

This extension adds a two small structures to memory requirements
querying and memory allocation: a new structure that flags whether an
image/buffer should have a dedicated allocation, and a structure
indicating the image or buffer that an allocation will be bound to.

## <a href="#_promotion_to_vulkan_1_1" class="anchor"></a>Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with
the KHR suffix omitted. The original type, enum and command names are
still available as aliases of the core functionality.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html):

  - [VkMemoryDedicatedAllocateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfoKHR.html)

- Extending [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements2.html):

  - [VkMemoryDedicatedRequirementsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedRequirementsKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_DEDICATED_ALLOCATION_EXTENSION_NAME`

- `VK_KHR_DEDICATED_ALLOCATION_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS_KHR`

## <a href="#_examples" class="anchor"></a>Examples

``` c
    // Create an image with a dedicated allocation based on the
    // implementation's preference

    VkImageCreateInfo imageCreateInfo =
    {
        // Image creation parameters
    };

    VkImage image;
    VkResult result = vkCreateImage(
        device,
        &imageCreateInfo,
        NULL,               // pAllocator
        &image);

    VkMemoryDedicatedRequirementsKHR dedicatedRequirements =
    {
        .sType = VK_STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS_KHR,
        .pNext = NULL,
    };

    VkMemoryRequirements2 memoryRequirements =
    {
        .sType = VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2,
        .pNext = &dedicatedRequirements,
    };

    const VkImageMemoryRequirementsInfo2 imageRequirementsInfo =
    {
        .sType = VK_STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2,
        .pNext = NULL,
        .image = image
    };

    vkGetImageMemoryRequirements2(
        device,
        &imageRequirementsInfo,
        &memoryRequirements);

    if (dedicatedRequirements.prefersDedicatedAllocation) {
        // Allocate memory with VkMemoryDedicatedAllocateInfoKHR::image
        // pointing to the image we are allocating the memory for

        VkMemoryDedicatedAllocateInfoKHR dedicatedInfo =
        {
            .sType = VK_STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO_KHR,
            .pNext = NULL,
            .image = image,
            .buffer = VK_NULL_HANDLE,
        };

        VkMemoryAllocateInfo memoryAllocateInfo =
        {
            .sType = VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO,
            .pNext = &dedicatedInfo,
            .allocationSize = memoryRequirements.size,
            .memoryTypeIndex = FindMemoryTypeIndex(memoryRequirements.memoryTypeBits),
        };

        VkDeviceMemory memory;
        vkAllocateMemory(
            device,
            &memoryAllocateInfo,
            NULL,               // pAllocator
            &memory);

        // Bind the image to the memory

        vkBindImageMemory(
            device,
            image,
            memory,
            0);
    } else {
        // Take the normal memory sub-allocation path
    }
```

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-02-27 (James Jones)

  - Copy content from VK_NV_dedicated_allocation

  - Add some references to external object interactions to the overview.

- Revision 2, 2017-03-27 (Faith Ekstrand)

  - Rework the extension to be query-based

- Revision 3, 2017-07-31 (Faith Ekstrand)

  - Clarify that memory objects allocated with
    VkMemoryDedicatedAllocateInfoKHR can only have the specified
    resource bound and no others.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_dedicated_allocation"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
